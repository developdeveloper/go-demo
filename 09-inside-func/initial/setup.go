package initial

import "fmt"

func init() {
	fmt.Println("init called in initial package")
}

func init() {
	fmt.Println("called in initial package too")
}

//DoNothing 啥也不干的函数
func DoNothing() {}
