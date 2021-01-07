package channel

import (
	"fmt"
	"sync"
	"time"
)

func inc(x *int, lock chan struct{}) {
	lock <- struct{}{}
	*x = *x + 1
	<-lock
}

func sumByLoop() {
	x := 0
	lock := make(chan struct{}, 1)
	for i := 0; i < 100; i++ {
		go inc(&x, lock)
	}
	// 为了代码结构简洁，简单的加个停顿(应该使用 sync.waitGroup 实现)
	time.Sleep(1 * time.Second)
	fmt.Println(x)
}

func inc2(x *int, mutex *sync.Mutex) {
	mutex.Lock()
	*x = *x + 1
	mutex.Unlock()
}

func sumByLoop2() {
	x := 0
	mutex := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		go inc2(&x, mutex)
	}
	// 为了代码结构简洁，简单的加个停顿(应该使用 sync.waitGroup 实现)
	time.Sleep(1 * time.Second)
	fmt.Println(x)
}
