package channel

import "fmt"

func lenCap() {
	d1 := make(chan int)
	d2 := make(chan int, 3)
	d2 <- 1
	fmt.Println(len(d1), cap(d1)) // 0 0
	fmt.Println(len(d2), cap(d2)) // 1 3
}
