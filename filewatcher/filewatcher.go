/**
 * @file    filewatcher.go
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
	"os"
	"time"
)

type fileWatcher struct {
	file    string
	call    CallFunc
	modTime time.Time
}

func AddFileWatcher(file string, callback CallFunc) error {
	fw := &fileWatcher{
		file: file,
		call: callback,
	}
	info, err := os.Stat(file)
	if err != nil {
		return err
	}
	fw.modTime = info.ModTime()
	watcherManager.Lock()
	watcherManager.fileWatchers[file] = fw
	watcherManager.Unlock()
	return nil
}

func RemoveFileWatcher(file string) {
	watcherManager.Lock()
	delete(watcherManager.fileWatchers, file)
	watcherManager.Unlock()
}
