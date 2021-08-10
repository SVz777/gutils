/**
 * @file    watcher.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020-03-02
 * @desc
 */
package filewatcher

import (
	"github.com/SVz777/gutils/collections"
	"github.com/SVz777/gutils/collections/set"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

type Op uint32

func (op Op) String() string {
	switch op {
	case Create:
		return "Create"
	case Modify:
		return "Modify"
	case Delete:
		return "Delete"
	default:
		return ""
	}
}

const (
	Create Op = 1 << iota
	Modify
	Delete
)

var ScanInterval = 2 * time.Second

type CallFunc func(file string, opType Op) error

type watcher struct {
	sync.RWMutex
	fileWatchers map[string]*fileWatcher
	dirWatchers  map[string]*dirWatcher
}

var watcherManager *watcher

func init() {
	watcherManager = &watcher{
		fileWatchers: make(map[string]*fileWatcher),
		dirWatchers:  make(map[string]*dirWatcher),
	}
	watcherManager.run()
}

func (w *watcher) run() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("file watcher fatal:", err)
			}
		}()
		// all file watchers
		for {
			select {
			case <-time.After(ScanInterval):
				w.RLock()
				for _, fw := range w.fileWatchers {
					info, err := os.Stat(fw.file)
					if err != nil {
						log.Println("os stat error:", err)
						continue
					}
					if fw.modTime.Before(info.ModTime()) {
						// 修改过
						if err := fw.call(fw.file, Modify); err != nil {
							log.Println("fw callback error:", err)
						}
						fw.modTime = info.ModTime()
					}
				}
				w.RUnlock()
			}
		}

	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("dir watcher fatal:", err)
			}
		}()
		// all dir watchers
		for {
			select {
			case <-time.After(ScanInterval):
				w.RLock()
				for _, dw := range w.dirWatchers {
					nowfiles, err := ioutil.ReadDir(dw.dir)
					if err != nil {
						log.Println("dw read dir error:", err)
						continue
					}
					dw.RWMutex.Lock()
					beforeFiles := set.Set{}
					for file := range dw.files {
						beforeFiles[file] = collections.Empty{}
					}

					for _, file := range nowfiles {
						if modTime, ok := dw.files[file.Name()]; ok {
							// 删除访问过的文件
							delete(beforeFiles, file.Name())
							if modTime.Before(file.ModTime()) {
								// 修改过
								if err := dw.call(file.Name(), Modify); err != nil {
									log.Println("dw callback error:", err)
									continue
								}
								dw.files[file.Name()] = file.ModTime()
							}
						} else {
							if err := dw.call(file.Name(), Create); err != nil {
								log.Println("dw callback error:", err)
								continue
							}
							dw.files[file.Name()] = file.ModTime()
						}

					}

					for file := range beforeFiles {
						// beforeFiles中都是没访问过的，也就是删除的
						if err := dw.call(file, Delete); err != nil {
							log.Println("dw callback error:", err)
							continue
						}
						delete(dw.files, file)
					}
					dw.RWMutex.Unlock()
				}
				w.RUnlock()
			}
		}
	}()
}
