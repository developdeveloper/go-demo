package main

import (
	"fmt"
	"unsafe"
)

// 目前的编译器，结构体成员是按顺序存放的
// 为了性能未来编译器可能会采用字节对齐，修改字段的顺序，就不能简单是使用 Offsetof 来操作了
type person struct {
	Name  string
	Money int // int8 int16 int64
}

func main() {
	p := person{Name: "zhangsan", Money: 100}
	fmt.Printf("memory size: %d\n", unsafe.Sizeof(p)) // 24
	fmt.Printf("align size: %d\n", unsafe.Alignof(p)) // 8

	fmt.Printf("%p\n", &p)
	pointer := unsafe.Pointer(&p) // 0xc0000a6040
	fmt.Println(*(*person)(pointer))

	namePtr := (*string)(pointer)
	*namePtr = "lisi"

	// 第一个成员的地址和结构体的地址一致
	// 0xc0000a6040
	fmt.Printf("%p\n", namePtr)

	//这种写法可以运行，但是被认为可能被 GC 回收导致不可预期的错误
	//TODO 重现的方法
	//ptr := uintptr(pointer)
	//agePtr := (*int64)(unsafe.Pointer(ptr + unsafe.Offsetof(p.Money)))

	agePtr := (*int64)(unsafe.Pointer(uintptr(pointer) + unsafe.Offsetof(p.Money)))
	*agePtr = *agePtr + 100

	// 0xc0000a6050
	// reflect.StringHeader(Data Len) 需要 16 个字节
	fmt.Printf("%p\n", agePtr)

	fmt.Printf("name = %s, money = %d", p.Name, p.Money)
}
