package main

import "fmt"

type person struct {
	Name string
}

/*
	0x001d 00029 (main.go:9)	LEAQ	type."".person(SB), AX
	0x0024 00036 (main.go:9)	MOVQ	AX, (SP)
	0x0028 00040 (main.go:9)	PCDATA	$1, $0
	0x0028 00040 (main.go:9)	CALL	runtime.newobject(SB)
	0x002d 00045 (main.go:9)	MOVQ	8(SP), AX
	0x0032 00050 (main.go:9)	MOVQ	$8, 8(AX)
	0x003a 00058 (main.go:9)	LEAQ	go.string."zhangsan"(SB), CX
	0x0041 00065 (main.go:9)	MOVQ	CX, (AX)
	0x0044 00068 (main.go:9)	MOVQ	AX, "".~r0+32(SP)
	0x0049 00073 (main.go:9)	MOVQ	16(SP), BP
	0x004e 00078 (main.go:9)	ADDQ	$24, SP
*/
func foo1() *person {
	// 分配到堆上
	return &person{Name: "zhangsan"}
}

/*
	0x0000 00000 (main.go:14)	LEAQ	go.string."zhangsan"(SB), AX
	0x0007 00007 (main.go:14)	MOVQ	AX, "".~r0+8(SP)
	0x000c 00012 (main.go:14)	MOVQ	$8, "".~r0+16(SP)
	0x0015 00021 (main.go:14)	RET
*/
func foo2() person {
	// 分配到栈上
	return person{Name: "zhangsan"}
}

/*
	0x0020 00032 (main.go:21)	CALL	runtime.printlock(SB)
	0x0025 00037 (main.go:21)	LEAQ	go.string."zhangsan"(SB), AX
	0x002c 00044 (main.go:21)	MOVQ	AX, (SP)
	0x0030 00048 (main.go:21)	MOVQ	$8, 8(SP)
	0x0039 00057 (main.go:21)	CALL	runtime.printstring(SB)
	0x003e 00062 (main.go:21)	NOP
	0x0040 00064 (main.go:21)	CALL	runtime.printnl(SB)
	0x0045 00069 (main.go:21)	CALL	runtime.printunlock(SB)
	0x004a 00074 (main.go:22)	MOVQ	16(SP), BP
	0x004f 00079 (main.go:22)	ADDQ	$24, SP
	0x0053 00083 (main.go:22)	RET
*/

// new(person) does not escape
func foo3() {
	// 分配到栈上
	p := new(person)
	p.Name = "zhangsan"
	println(p)
}

// new(person) escapes to heap
func foo3_1() {
	// 分配到堆上
	p := new(person)
	p.Name = "zhangsan"
	// 传递的是 []interface{}
	fmt.Printf("%v", p)
}

/*
	0x001d 00029 (main.go:26)	LEAQ	type."".person(SB), AX
	0x0024 00036 (main.go:26)	MOVQ	AX, (SP)
	0x0028 00040 (main.go:26)	PCDATA	$1, $0
	0x0028 00040 (main.go:26)	CALL	runtime.newobject(SB)
	0x002d 00045 (main.go:26)	MOVQ	8(SP), AX
	0x0032 00050 (main.go:27)	MOVQ	$8, 8(AX)
	0x003a 00058 (main.go:27)	LEAQ	go.string."zhangsan"(SB), CX
	0x0041 00065 (main.go:27)	MOVQ	CX, (AX)
	0x0044 00068 (main.go:28)	MOVQ	AX, "".~r0+32(SP)
	0x0049 00073 (main.go:28)	MOVQ	16(SP), BP
	0x004e 00078 (main.go:28)	ADDQ	$24, SP
	0x0052 00082 (main.go:28)	RET
*/

// new(person) escapes to heap
func foo4() *person {
	// 分配到堆上
	p := new(person)
	p.Name = "zhangsan"
	return p
}
