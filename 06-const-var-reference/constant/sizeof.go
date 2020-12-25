package constant

import (
	"fmt"
	"unsafe"
)

// CompileSizedof 编译器确定所以可以给 const 赋值
func CompileSizedof() {
	// str := "english 中文"
	var str string = "english 中文"
	const byteInSize = unsafe.Sizeof(str)
	// 不管 str 是什么内容 byteInSize 都是 16
	// 字符串在 go 内部用 2 个字段表示(一个地址值和一个长度值，都占用 8 个字节)
	fmt.Println(byteInSize)
}
