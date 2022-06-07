/**
 * @file    multitask.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020/11/10
 * @desc
 */
package multitask

import (
	"context"
	"fmt"
	"sync"
)

type TaskManager struct {
	ctx   context.Context
	wg    sync.WaitGroup
	tasks map[string]ITask
}

func NewTaskManager(ctx context.Context) *TaskManager {
	return &TaskManager{
		ctx:   ctx,
		wg:    sync.WaitGroup{},
		tasks: make(map[string]ITask),
	}
}

func (tm *TaskManager) Add(key string, f Do) {
	tm.tasks[key] = NewTask(tm.ctx, key, f)
}

func (tm *TaskManager) GetAllTasks() map[string]ITask {
	return tm.tasks
}

func (tm *TaskManager) GetTaskResult(key string) (result interface{}) {
	if v, ok := tm.tasks[key]; ok {
		result, _ = v.GetResult()
	}
	return
}

func (tm *TaskManager) GetTaskErr(key string) (err error) {
	if v, ok := tm.tasks[key]; ok {
		_, err = v.GetResult()
	}
	return
}

func (tm *TaskManager) Do() error {
	if len(tm.tasks) <= 0 {
		return nil
	}
	tm.wg.Add(len(tm.tasks))
	for key, t := range tm.tasks {
		go func(ctx context.Context, key string, t ITask) {
			defer tm.wg.Done()
			go t.Do(0)
			select {
			case <-t.Done():
				return
			case <-ctx.Done():
				t.SetResult(nil, fmt.Errorf("ctx err: %w", ctx.Err()))
				return
			}
		}(tm.ctx, key, t)
	}
	tm.wg.Wait()
	var err error
	for key, task := range tm.tasks {
		_, taskErr := task.GetResult()
		if taskErr != nil {
			if err == nil {
				err = fmt.Errorf("[%v:%v]", key, taskErr)
			} else {
				err = fmt.Errorf("[%v:%v]||%w", key, taskErr, err)
			}
		}
	}
	return err
}
