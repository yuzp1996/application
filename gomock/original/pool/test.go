package pool

import (
	"github.com/stretchr/testify/assert"
	"github.com/yuzp1996/application/gomock/original/mock"
	"github.com/yuzp1996/application/gomock/original/task"
	"testing"
)

type TestSuit struct {
	name    string
	newTask NewTask
	size    int
	times   int
}

func TestTaskPoolRunImpl(t *testing.T) {

	testSuits := []TestSuit{
		{
			name:    "minimal task pool",
			newTask: func() task.Task { return mock.NewMinimalTask() },
			size:    100,
			times:   200,
		},
	}

	for _, suit := range testSuits {
		t.Run(suit.name, func(t *testing.T) {
			var taskPool TaskPool = NewTaskPoolImpl(suit.newTask, suit.size)
			err := taskPool.Run(suit.size)
			assert.NoError(t, err)
		})
	}
}
