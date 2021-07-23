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

const defaultTaskNum = 4 // 默认任务数，大部分并行任务<=4个

type do func(context.Context) (interface{}, error)

type task struct {
	f      do
	result interface{}
	err    error
}

func (task *task) GetResult() (ret interface{}) {
	return task.result
}

func (task *task) GetErr() (ret error) {
	return task.err
}

type TaskManager struct {
	ctx   context.Context
	wg    sync.WaitGroup
	tasks map[string]*task
}

func NewTaskManager(ctx context.Context) *TaskManager {
	return &TaskManager{
		ctx:   ctx,
		wg:    sync.WaitGroup{},
		tasks: make(map[string]*task, defaultTaskNum),
	}
}

func (tm *TaskManager) Add(key string, f do) {
	tm.tasks[key] = &task{f: f}
}

func (tm *TaskManager) GetAllTasks() map[string]*task {
	return tm.tasks
}

func (tm *TaskManager) GetTaskResult(key string) (result interface{}) {
	if v, ok := tm.tasks[key]; ok {
		result = v.result
	}
	return
}

func (tm *TaskManager) GetTaskErr(key string) (err error) {
	if v, ok := tm.tasks[key]; ok {
		return v.err
	}
	return
}

func (tm *TaskManager) Do() error {
	if len(tm.tasks) <= 0 {
		return nil
	}
	tm.wg.Add(len(tm.tasks))
	for key, t := range tm.tasks {
		go func(ctx context.Context, key string, t *task) {
			defer func() {
				if p := recover(); p != nil {
					t.err = fmt.Errorf("task %v err: %v", key, p)
				}
			}()
			defer tm.wg.Done()
			t.result, t.err = t.f(ctx)
		}(tm.ctx, key, t)
	}
	tm.wg.Wait()
	var err error
	for key, task := range tm.tasks {
		if task.err != nil {
			if err == nil {
				err = fmt.Errorf("[%v:%v]", key, task.err)
			} else {
				err = fmt.Errorf("[%v:%v]||%w", key, task.err, err)
			}
		}
	}
	return err
}
