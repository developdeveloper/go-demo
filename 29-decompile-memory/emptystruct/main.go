package main

import "fmt"

func main() {
	s := &struct{}{}
	fmt.Println(*s)
}

/*
0x001d 00029 (main.go:7)	XORPS	X0, X0                    // 清空 X0 寄存器
0x0020 00032 (main.go:7)	MOVUPS	X0, ""..autotmp_15+64(SP) // 初始化栈上临时空间
0x0025 00037 (main.go:7)	LEAQ	type.struct {}(SB), AX
0x002c 00044 (main.go:7)	MOVQ	AX, ""..autotmp_15+64(SP)
0x0031 00049 (main.go:7)	LEAQ	runtime.zerobase(SB), AX
0x0038 00056 (main.go:7)	MOVQ	AX, ""..autotmp_15+72(SP)
*/
