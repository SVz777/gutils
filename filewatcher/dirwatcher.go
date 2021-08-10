/**
 * @file    dirwatcher.go
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
	"io/ioutil"
	"sync"
	"time"
)

type dirWatcher struct {
	sync.RWMutex
	dir   string
	call  CallFunc
	files map[string]time.Time
}

func AddDirWatcher(path string, callback CallFunc) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	dw := &dirWatcher{
		dir:   path,
		call:  callback,
		files: make(map[string]time.Time, len(files)),
	}
	for _, file := range files {
		dw.files[file.Name()] = file.ModTime()
	}
	watcherManager.Lock()
	watcherManager.dirWatchers[path] = dw
	watcherManager.Unlock()
	return nil
}

func RemoveDirWatcher(path string) {
	watcherManager.Lock()
	delete(watcherManager.dirWatchers, path)
	watcherManager.Unlock()
}
