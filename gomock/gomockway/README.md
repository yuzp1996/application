 
 ```cassandraql
 mockgen -destination gomockway/mock_task.go -package mock -source original/task/task.go
```
 生成对应的mock文件  这样就可以去调用了
 
 
 ```cassandraql
关键代码：

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()


    newTask: func() task.Task {
        mockTask := mock.NewMockTask(ctrl)
        //定义了置入的值和返回的值
        mockTask.EXPECT().Do(gomock.Any()).Return("", errors.New("return err")).AnyTimes()
        return mockTask
    },


	for _, suit := range testSuits {
		t.Run(suit.name, func(t *testing.T) {
			taskPool = NewTaskPoolImpl(suit.newTask, suit.size)
            // 这里会调用上面中的Do方法 而且将值返回来
			err := taskPool.Run(suit.times)
			if suit.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}




```
 
 mock的意义
 
 1. 觉得写mock的是约定好接口，然后在面向接口做开发的时候能够方便测试，因为不需要接口实际的实现，而是依赖mock的Minimal Implement就可以进行单元测试
 2. 应该对接口返回什么值敏感，知道接口返回什么值  会对自己的程序造成什么样的影响  