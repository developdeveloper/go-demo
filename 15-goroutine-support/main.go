package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	go func() {
// 		fmt.Println("我来自协程")
// 	}()

// 	fmt.Println("wait...")

// 	// select {}
// 	time.Sleep(3 * time.Second)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go func(index int) {
// 			fmt.Printf("我来自协程 %d\n", index)
// 			time.Sleep(1 * time.Second)
// 		}(i)
// 	}

// 	fmt.Println("wait...")

// 	// select {}
// 	time.Sleep(3 * time.Second)
// }

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 对协程进行计数

		go func(index int) {
			fmt.Printf("我来自协程 %d\n", index)
			wg.Done() // 通知一下，本协程把问题搞完了
		}(i)
	}

	fmt.Println("wait...")
	wg.Wait() // 等待所有的协程通知我
}
