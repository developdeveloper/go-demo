package locker

import (
	"fmt"
	"sync"
)

func lockFailed() {
	var mu sync.Mutex
	go func() {
		fmt.Println("NO.1中国人")
		mu.Lock()
	}()
	mu.Unlock()
}
