package channel

import "fmt"

func forRangeChan() {
	queue := make(chan string, 3)

	go func() {
		queue <- "one"
		queue <- "two"
		queue <- "three"
		close(queue) // 去掉 close for-range 将死锁
	}()

	for elem := range queue {
		fmt.Println(elem)
	}
}
