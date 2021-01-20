package main

import (
	"fmt"
	"time"
)

func main() {
	num := 0

	go func() {
		num = 1
	}()

	num = 2
	fmt.Println(num)
	time.Sleep(3 * time.Second)
}
