package iface

import "fmt"

func emptyInterface() {
	var any interface{} = 0 // 给空接口对象初值，否则默认是 nil
	fmt.Printf("%T\n", any) // int

	any = "test"
	fmt.Printf("%T\n", any) // string

	any = 3.1415926
	fmt.Printf("%T\n", any) // float64

	any = []int{1, 2, 3}
	fmt.Printf("%T\n", any) // []int

	any = struct{}{}
	fmt.Printf("%T\n", any) // struct {}
}
