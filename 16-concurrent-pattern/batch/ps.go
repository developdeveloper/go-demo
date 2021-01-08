package batch

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func producer(name string, tasks chan<- string) {
	index := 0
	for {
		tasks <- fmt.Sprintf("%s-task-%d", name, index)
		time.Sleep(50 * time.Millisecond)
		index++
	}
}

func producer2(max int, tasks chan<- string, doSth func() string) {
	lines := make(chan int, max)
	for {
		lines <- 1
		tasks <- doSth()
		<-lines
	}
}

func consumer(name string, tasks <-chan string) {
	for task := range tasks {
		fmt.Printf("%s-%s has been token\n", name, task)
		time.Sleep(50 * time.Millisecond)
	}
}

func createFactory(stockSize, producerCount, consumerCount int) {
	tasks := make(chan string, stockSize)
	for i := 0; i <= producerCount; i++ {
		go producer(fmt.Sprintf("p%d", i), tasks)
	}

	for i := 0; i <= consumerCount; i++ {
		go consumer(fmt.Sprintf("c%d", i), tasks)
	}

	exit := make(chan os.Signal, 1)
	// ctrl+c 中断程序
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
