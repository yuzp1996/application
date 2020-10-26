package pool

import (
	"fmt"

	"github.com/yuzp1996/application/gomock/original/task"
)

type TaskPoolImpl struct {
	pool []task.Task
}

func NewTaskPoolImpl(newTask NewTask, size int) *TaskPoolImpl {
	tp := &TaskPoolImpl{
		pool: make([]task.Task, size),
	}
	for i := 0; i < size; i++ {
		tp.pool[i] = newTask()
	}
	return tp
}

func (tp *TaskPoolImpl) Run(times int) error {
	poolLen := len(tp.pool)
	for i := 0; i < times; i++ {
		ret, err := tp.pool[i%poolLen].Do(i)
		if err != nil {
			// process error
			return err
		}
		switch ret {
		case "":
			// process 0
			fmt.Println(ret)
		case "a":
			// process 1
			fmt.Println(ret)
		case "b":
			// process 2
			fmt.Println(ret)
		case "c":
			// process 3
			fmt.Println(ret)
		}
	}
	return nil
}
