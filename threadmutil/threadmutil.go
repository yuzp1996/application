package threadmutil

import (
	"fmt"
	"sync"
)

type Truck struct {
	SUM      int         `json:"sum"`
	Capacity int         `json:"capacity"`
	Start    chan bool   `json:"start"`
	Last     chan bool   `json:"last"`
	Lock     sync.Locker `json:"lock"`
}

func (truck *Truck) Load(p Package) {
	go func() {
		fmt.Printf("Waiting for package our Pack %d \n", p.Identifier)

		truck.Lock.Lock()
		truck.SUM += 1
		// 下达装货指令
		if truck.SUM >= truck.Capacity {
			truck.Start <- true
		}
		truck.Lock.Unlock()

		// 阻塞线程  直到有东西开始读取
		truck.Last <- false
		fmt.Printf("Packing the package %d \n", p.Identifier)

		// 读取
		truck.Lock.Lock()
		truck.SUM -= 1
		if truck.SUM == 0 {
			truck.Last <- true
		}
		truck.Lock.Unlock()
	}()

}

type Package struct {
	Identifier int
}

// we have 100 packages when it is 101, we will load it
func TransformPackage() {
	start := make(chan bool)
	last := make(chan bool)

	truck := Truck{
		SUM:      0,
		Capacity: 100,
		Start:    start,
		Lock:     &sync.Mutex{},
		Last:     last,
	}

	// 开始准备货物
	for i := 1; i <= truck.Capacity; i++ {
		truck.Load(Package{Identifier: i})
	}
	// 等待开始装货命令
	<-truck.Start

	// 等待装货结束指令  通过 channel 来同步信息
	for last := range truck.Last {
		if last != true {
			// 这里总会得到一个值  就是读取值的没有互斥锁的原因
			fmt.Printf("Not finished, truck sum is  %d\n", truck.SUM)
		} else {
			fmt.Printf("Finished!!!")
			break
		}
	}

}
