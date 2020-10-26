package mock

import "github.com/golang/mock/gomock"

func NewMockMinimalTask(ctrl *gomock.Controller, num int) *MockTask {
	mock := NewMockTask(ctrl)
	mock.EXPECT().Do(num).Return("", nil).AnyTimes()
	return mock
}
