package timelib

import (
	"fmt"
	"time"
)

//TimerAfterTest 计时器
func TimerAfterTest() {
	signal := time.After(1 * time.Second)
	select {
	case <-signal:
		fmt.Println("到时间了...")
	case <-time.After(2 * time.Second):
		fmt.Println("超过2秒了...")
	}

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("我1秒后打印")
	})
}

// TickerTest 定时器
func TickerTest() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case <-ticker.C:
			fmt.Println("时间它滴答滴答滴")
		}
	}
}

//TickTest Tick 函数返回的通道，没有关闭方法，一直用
func TickTest() {
	tickChan := time.Tick(1 * time.Second)

	go func() {
		for {
			select {
			case <-tickChan:
				fmt.Println("进程结束前我一直打印")
			}
		}
	}()
}
