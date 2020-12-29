package scope

import (
	"fmt"
	"time"
)

func createPanic() {
	defer func() {
		fmt.Println("i will not be tiggered")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("will panic?")
	go func() {
		defer fmt.Println("i will be tiggered")
		panic("can you handle me?")
	}()

	// 等一会儿结束主线程
	time.Sleep(3 * time.Second)
}
