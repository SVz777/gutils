/**
 * @file    task.go
 * @author  903943711@qq.com
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * @date    2021/8/24
 * @desc
 */
package multitask

import (
	"context"
	"log"
	"sync"

	"github.com/SVz777/gutils/collections"
)

type Do func(context.Context) (interface{}, error)

type ITask interface {
	// Do pool core id执行task
	Do(id int)
	// Context 获取task的context
	Context() context.Context
	// GetKey 获取task key
	GetKey() string
	// SetResult 设置task 结果
	SetResult(interface{}, error)
	// GetResult 获取task 结果
	GetResult() (interface{}, error)
	// Done 返回一个完成标记的chan
	Done() <-chan collections.Empty
}

type Task struct {
	ctx context.Context
	key string
	f   Do

	runID  int
	result interface{}
	err    error
	done   chan collections.Empty
	once   sync.Once
}

func NewTask(ctx context.Context, key string, f Do) *Task {
	return &Task{ctx: ctx, key: key, f: f, done: make(chan collections.Empty, 1)}
}

func (task *Task) Context() context.Context {
	return task.ctx
}
func (task *Task) Do(id int) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("process: %v task: %v err: %v\n", id, task.GetKey(), err)
		}
	}()
	task.runID = id
	task.SetResult(task.f(task.ctx))
	task.end()
}

func (task *Task) SetResult(result interface{}, err error) {
	task.once.Do(func() {
		if err != nil {
			task.err = err
		} else {
			task.result, task.err = result, err
		}
	})
}

func (task *Task) GetKey() string {
	return task.key
}

func (task *Task) GetResult() (interface{}, error) {
	return task.result, task.err
}

func (task *Task) Done() <-chan collections.Empty {
	return task.done
}

func (task *Task) end() {
	close(task.done)
}
