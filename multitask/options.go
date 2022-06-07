package multitask

type Options struct {
	TaskNum int
}
type Option func(o *Options)

func WithTaskNum(num int) Option {
	return func(o *Options) {
		o.TaskNum = num
	}
}
