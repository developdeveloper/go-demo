package batch

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func handle(tasks chan string, index int) {
	defer wg.Done()
	for {
		// 读取关闭的无缓冲通道，返回零值和false
		// 读取关闭的有缓冲通道，先把缓存数据读完后，再读取返回零值和false
		task, ok := <-tasks
		if !ok {
			// 任务通道关闭了，不会再有新任务了
			fmt.Printf("handle %d over\n", index)
			return
		}
		fmt.Printf("干活儿 %s By 工人 %d\n", task, index)
	}
}

func createTasks(goroutines, taskCount int) {
	tasks := make(chan string, taskCount)

	// 启动 N
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go handle(tasks, i)
	}

	// 设置 M
	for k := 0; k <= taskCount; k++ {
		tasks <- fmt.Sprintf("任务 %d", k)
	}

	close(tasks)
	wg.Wait()
}
