package delay

import "fmt"

// 全局
var globalFlag int = 0

func flagFunc() int {
	defer func() {
		globalFlag = 1
	}()

	fmt.Printf("callAfterReturn-flag = %d\n", globalFlag) // 0
	return globalFlag
}

func callAfterReturn() {
	localFlag := flagFunc()
	fmt.Printf("callAfterReturn-localFlag = %d\n", localFlag)   // 0
	fmt.Printf("callAfterReturn-globalFlag = %d\n", globalFlag) // 1
}
