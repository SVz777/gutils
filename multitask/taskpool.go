/**
 * @file    taskpool.go
 * @author  903943711@qq.com
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * @date    2021/8/23
 * @desc
 */
package multitask

type ITaskPool interface {
	Do(ITask)
}

func NewTaskPool(opt ...Option) ITaskPool {
	opts := GetDefaultOptions()
	for _, o := range opt {
		o(opts)
	}
	if opts.HashFunc == nil {
		return newNormalTaskPool(opts)
	}
	return newHashTaskPool(opts)
}

// normalTaskPool 随机分配core执行task
type normalTaskPool struct {
	ch   chan ITask
	core int
}

func newNormalTaskPool(opts *Options) *normalTaskPool {
	tp := normalTaskPool{
		ch:   make(chan ITask, opts.CoreNum),
		core: opts.CoreNum,
	}
	for ii := 0; ii < tp.core; ii++ {
		go start(ii, tp.ch)
	}
	return &tp
}

func (tp *normalTaskPool) Do(task ITask) {
	tp.ch <- task
}

// hashTaskPool 同key task会在同一个core上执行，保证有序
type hashTaskPool struct {
	chs  []chan ITask
	core int
	hash func(string) int
}

func newHashTaskPool(opts *Options) *hashTaskPool {
	tp := hashTaskPool{
		chs:  make([]chan ITask, opts.CoreNum),
		core: opts.CoreNum,
		hash: opts.HashFunc,
	}
	for idx := range tp.chs {
		tp.chs[idx] = make(chan ITask, opts.CoreNum)
	}
	for ii := 0; ii < tp.core; ii++ {
		go start(ii, tp.chs[ii])
	}
	return &tp
}

func (tp *hashTaskPool) Do(task ITask) {
	tp.chs[tp.hash(task.GetKey())%tp.core] <- task
}

func start(id int, c chan ITask) {
	for {
		task := <-c
		select {
		case <-task.Context().Done():
			// context 结束了不执行task
			continue
		default:
			task.Do(id)
		}
	}
}
