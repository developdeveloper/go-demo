package locker

import "fmt"

func lockSuccess() {
	exit := make(chan int)
	go func() {
		fmt.Println("NO.1中国人")
		<-exit
	}()
	exit <- 1
}

func lockSuccess2() {
	exit := make(chan struct{})
	go func() {
		fmt.Println("NO.1中国人")
		exit <- struct{}{}
	}()
	<-exit
}

func lockBatch(taskCount int) {
	waitGroup := make(chan struct{}, taskCount)
	for i := 0; i < cap(waitGroup); i++ {
		go func(index int) {
			fmt.Printf("NO.1中国人%d\n", index)
			waitGroup <- struct{}{}
		}(i)
	}
	for i := 0; i < cap(waitGroup); i++ {
		<-waitGroup
	}
}
