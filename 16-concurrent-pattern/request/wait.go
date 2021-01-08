package request

import (
	"fmt"
	"time"
)

type request struct {
	data   []interface{} // 模拟要处理的数据
	finish chan struct{} // 完成后的标志
}

func newRequest(data ...interface{}) *request {
	return &request{data, make(chan struct{}, 1)}
}

func doWork(req *request) {
	time.Sleep(1 * time.Second) // 模拟耗时
	fmt.Println(req.data)       // 工作内容就是打印
	close(req.finish)           // req.finish <- struct{}{} 也行
}

func createWork() {
	req := newRequest(1, "string", true)
	go doWork(req)

	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("继续做事，我很忙")
	}

	<-req.finish
}
