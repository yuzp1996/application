package pool

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yuzp1996/application/gomock/gomockway/mock"
	"github.com/yuzp1996/application/gomock/gomockway/task"
	"testing"
)

type TestSuit struct {
	name    string
	newTask NewTask
	size    int
	times   int
	isErr   bool
}

func TestTaskPoolRunImpl_MinimalTask(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	strs := []string{"a", "b", "c"}
	count := 0
	size := 3
	rounds := 1

	testSuits := []TestSuit{
		//{
		//    name:    "minimal task pool",
		//    newTask: func() task.Task { return mock.NewMinimalTask() },
		//    size:    100,
		//    times:   200,
		//},
		{
			name:    "mock minimal task pool",
			newTask: func() task.Task { return mock.NewMockMinimalTask(ctrl, count) },
			size:    100,
			times:   200,
		},
		{
			name: "return err",
			newTask: func() task.Task {
				mockTask := mock.NewMockTask(ctrl)
				mockTask.EXPECT().Do(gomock.Any()).Return("", errors.New("return err")).AnyTimes()
				return mockTask
			},
			size:  100,
			times: 200,
			isErr: true,
		},
		{
			name: "check input and output",
			newTask: func() task.Task {
				mockTask := mock.NewMockTask(ctrl)
				// 这里我们通过Do的设置检查了mackTask.Do调用时候的入参以及调用次数
				// 通过Return来设置发生调用时的返回值
				mockTask.EXPECT().Do(count).Return(strs[count%3], nil).Times(rounds)
				count++
				return mockTask
			},
			size:  size,
			times: size * rounds,
			isErr: false,
		},
	}
	var taskPool TaskPool
	for _, suit := range testSuits {
		t.Run(suit.name, func(t *testing.T) {
			taskPool = NewTaskPoolImpl(suit.newTask, suit.size)
			err := taskPool.Run(suit.times)
			if suit.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}
