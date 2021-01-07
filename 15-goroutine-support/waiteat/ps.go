package waiteat

import (
	"fmt"
	"time"
)

func produce(ready chan bool) {
	fmt.Println("做饭中...")
	time.Sleep(50 * time.Microsecond)
	ready <- true
	// close(ready)
}

func produce2(ready chan bool) {
	fmt.Println("做饭中...")
	// 做饭成功的与否的算法
	success := false
	if time.Now().Second()%2 == 0 {
		success = true
	}
	ready <- success
	// close(ready) 关闭 success 为 false
}

func eatTest() {
	ready := make(chan bool) // 无缓冲
	go produce(ready)        // 去掉 go 死锁
	for {
		select {
		case <-ready:
			fmt.Println("可以吃了")
			return
		default:
			fmt.Println("我先干点别的事")
		}
	}
}

func eatTest2() {
	ready := make(chan bool) // 无缓冲
	go produce(ready)        // 去掉 go 死锁
	<-ready                  // 同步交接
	fmt.Println("可以吃了")
}

func eatTest3() {
	ready := make(chan bool) // 无缓冲
	go produce2(ready)       // 去掉 go 死锁
	success := <-ready       // 同步交接
	//success, ok := <-ready // 关闭的通道 ok 为 false
	if success {
		fmt.Println("可以吃了")
	} else {
		fmt.Println("喝西北风")
	}
}
