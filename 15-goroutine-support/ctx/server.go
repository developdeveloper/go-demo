package ctx

import (
	"context"
	"fmt"
	"time"
)

func handle(ctx context.Context, index int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("handle %d over", index)
			return
		default:
			time.Sleep(50 * time.Microsecond)
			fmt.Printf("干活儿 %d\n", index)
		}
	}
}

func contextServerTest() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		go handle(ctx, i)
	}
	time.Sleep(1 * time.Second)
	cancel() // 可以 defer cancel
	time.Sleep(1 * time.Second)
}

func contextServerTest2() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//cancel() // 仍然可以手动提前结束
	for i := 0; i < 3; i++ {
		go handle(ctx, i)
	}
	time.Sleep(2 * time.Second)
}

func contextServerTest3() {
	parent, parentCancel := context.WithCancel(context.Background())
	ctx1, _ := context.WithCancel(parent)
	ctx2, _ := context.WithTimeout(parent, 10*time.Second)
	for i := 0; i < 3; i++ {
		go handle(ctx1, i)
	}
	for i := 3; i < 6; i++ {
		go handle(ctx2, i)
	}
	time.Sleep(1 * time.Second)
	parentCancel() // 子 ctx1,ctx2 也会遭殃
	time.Sleep(1 * time.Second)
}
