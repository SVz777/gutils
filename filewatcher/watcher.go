package filewatcher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/SVz777/gutils/collections"
	"github.com/SVz777/gutils/multitask"
)

type Watcher struct {
	sync.RWMutex
	fileWatchers map[string]*fileWatcher
	dirWatchers  map[string]*dirWatcher
	closed       chan collections.Empty
	opts         *Options
	tp           multitask.ITaskPool
}

func NewWatcher(opt ...Option) *Watcher {
	opts := GetDefaultOptions()
	opts.Update(opt...)
	return &Watcher{
		fileWatchers: make(map[string]*fileWatcher),
		dirWatchers:  make(map[string]*dirWatcher),
		closed:       make(chan collections.Empty),
		opts:         opts,
		tp:           multitask.NewTaskPool(opts.TpOption...),
	}
}

func (w *Watcher) Run(ctx context.Context) {
	tc := time.NewTicker(w.opts.ScanInterval)
	defer tc.Stop()
	for {
		select {
		case <-tc.C:
			for _, fw := range w.fileWatchers {
				tfw := fw
				w.tp.Do(multitask.NewTask(ctx, tfw.path, func(ctx context.Context) (interface{}, error) {
					fmt.Println(tfw.path, "run")
					tfw.watch()
					return nil, nil
				}))
			}
			for _, dw := range w.dirWatchers {
				tdw := dw
				w.tp.Do(multitask.NewTask(ctx, tdw.path, func(ctx context.Context) (interface{}, error) {
					fmt.Println(tdw.path, "run")
					tdw.watch()
					return nil, nil
				}))
			}
		}

	}

}

func (w *Watcher) AddFileWatcher(file string, callback CallFunc) error {
	fw, err := NewFileWatcher(file, callback)
	if err != nil {
		return err
	}
	w.Lock()
	w.fileWatchers[file] = fw
	w.Unlock()
	return nil
}

func (w *Watcher) RemoveFileWatcher(file string) {
	w.Lock()
	delete(w.fileWatchers, file)
	w.Unlock()
}

func (w *Watcher) AddDirWatcher(path string, callback CallFunc) error {
	dw, err := NewDirWatcher(path, callback)
	if err != nil {
		return err
	}
	w.Lock()
	w.dirWatchers[path] = dw
	w.Unlock()
	return nil
}

func (w *Watcher) RemoveDirWatcher(path string) {
	w.Lock()
	delete(w.dirWatchers, path)
	w.Unlock()
}

func (w *Watcher) Stop() {
	close(w.closed)
}
