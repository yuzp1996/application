package pool

import (
	"github.com/yuzp1996/application/gomock/gomockway/task"
)

type TaskPool interface {
	Run(times int) error
}

type NewTask func() task.Task
