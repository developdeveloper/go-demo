package channel

import (
	"fmt"
	"time"
)

func createBufChanVar() {
	//channel := make(chan string)
	channel := make(chan string, 1)
	go func() {
		fmt.Println(<-channel)
	}()

	channel <- "我来啦"

	// 死锁
	//channel <- "我来啦"
	//channel <- "我来啦"
	//fatal error: all goroutines are asleep - deadlock!

	for i := 0; i < 10; i++ {
		fmt.Printf("for 循环 %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
