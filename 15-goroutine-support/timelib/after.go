package timelib

import (
	"fmt"
	"time"
)

//timeAfterTest 时间等待
func timeAfterTest() {
	deadline := time.After(1 * time.Second)
	fmt.Printf("%T\n", deadline) // <-chan time.Time

	select {
	case <-deadline:
		fmt.Println("到时间了")
	}
}
