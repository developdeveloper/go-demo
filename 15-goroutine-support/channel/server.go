package channel

import (
	"fmt"
	"sync"
	"time"
)

//server 服务实例
type server struct {
	stopFlag  chan struct{}
	waitGroup sync.WaitGroup
}

//Work 做事
func (s *server) Work(index int) {
	time.Sleep(50 * time.Microsecond)
	fmt.Printf("干活儿 %d\n", index)
}

//Stop 停止
func (s *server) Stop() {
	close(s.stopFlag)  // 直接关闭通道
	s.waitGroup.Wait() // 等带 handle 的 goroutine 都正常结束
}

func (s *server) handle(index int) {
	s.waitGroup.Add(1)
	go func() {
		for {
			select {
			case <-s.stopFlag:
				fmt.Printf("handle %d over", index)
				s.waitGroup.Done()
				return
			default:
				s.Work(index)
			}
		}
	}()
}

func newServer() *server {
	return &server{make(chan struct{}), sync.WaitGroup{}}
}

func createServerTest() {
	srv := newServer()
	for i := 0; i < 3; i++ {
		srv.handle(i)
	}
	select {
	case <-time.After(500 * time.Microsecond):
		srv.Stop()
	}
}
