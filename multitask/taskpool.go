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

// NormalTaskPool 随机分配core执行task
type NormalTaskPool struct {
	ch   chan ITask
	core int
}

func NewNormalTaskPool(num int) *NormalTaskPool {
	tp := NormalTaskPool{
		ch:   make(chan ITask, num),
		core: num,
	}
	for ii := 0; ii < tp.core; ii++ {
		go start(ii, tp.ch)
	}
	return &tp
}

func (tp *NormalTaskPool) Do(task ITask) {
	tp.ch <- task
}

// HashTaskPool 同key task会在同一个core上执行，保证有序
type HashTaskPool struct {
	chs  []chan ITask
	core int
	hash func(string) int
}

func NewHashTaskPool(num int, hashFunc func(string) int) *HashTaskPool {
	tp := HashTaskPool{
		chs:  make([]chan ITask, num),
		core: num,
		hash: hashFunc,
	}
	for idx := range tp.chs {
		tp.chs[idx] = make(chan ITask, num)
	}
	for ii := 0; ii < tp.core; ii++ {
		go start(ii, tp.chs[ii])
	}
	return &tp
}

func (tp *HashTaskPool) Do(task ITask) {
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
