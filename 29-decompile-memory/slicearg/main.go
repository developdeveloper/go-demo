package main

import "fmt"

//go:noinline
func printSlice(slice []int) {
	fmt.Println(slice)
}

func main() {
	printSlice([]int{1, 2, 3})
}

/*
0x001d 00029 (main.go:11)	LEAQ	type.[3]int(SB), AX
0x0024 00036 (main.go:11)	MOVQ	AX, (SP)
0x0028 00040 (main.go:11)	PCDATA	$1, $0
0x0028 00040 (main.go:11)	CALL	runtime.newobject(SB)
0x002d 00045 (main.go:11)	MOVQ	8(SP), AX
0x0032 00050 (main.go:11)	MOVQ	$1, (AX)    // 元素1  8(SP)
0x0039 00057 (main.go:11)	MOVQ	$2, 8(AX)   // 元素2 16(SP)
0x0041 00065 (main.go:11)	MOVQ	$3, 16(AX)  // 元素3 24(SP)
0x0049 00073 (main.go:11)	MOVQ	AX, (SP)    // Data 地址
0x004d 00077 (main.go:11)	MOVQ	$3, 8(SP)   // Len 大小
0x0056 00086 (main.go:11)	MOVQ	$3, 16(SP)  // Cap 大小
0x005f 00095 (main.go:11)	NOP
0x0060 00096 (main.go:11)	CALL	"".printSlice(SB)
0x0065 00101 (main.go:12)	MOVQ	24(SP), BP
0x006a 00106 (main.go:12)	ADDQ	$32, SP
0x006e 00110 (main.go:12)	RET

"".printSlice STEXT size=175 args=0x18 locals=0x58
  0x0000 00000 (main.go:6)	TEXT	"".printSlice(SB), ABIInternal, $88-24
  0x0000 00000 (main.go:6)	MOVQ	(TLS), CX
  0x0009 00009 (main.go:6)	CMPQ	SP, 16(CX)
  0x000d 00013 (main.go:6)	PCDATA	$0, $-2
  0x000d 00013 (main.go:6)	JLS	165
  0x0013 00019 (main.go:6)	PCDATA	$0, $-1
  0x0013 00019 (main.go:6)	SUBQ	$88, SP
  0x0017 00023 (main.go:6)	MOVQ	BP, 80(SP)             // 基址
  0x001c 00028 (main.go:6)	LEAQ	80(SP), BP
  0x0021 00033 (main.go:6)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
  0x0021 00033 (main.go:6)	FUNCDATA	$1, gclocals·2589ca35330fc0fce83503f4569854a0(SB)
  0x0021 00033 (main.go:6)	FUNCDATA	$3, "".printSlice.stkobj(SB)
  0x0021 00033 (main.go:7)	MOVQ	"".slice+96(SP), AX    // *
  0x0026 00038 (main.go:7)	MOVQ	AX, (SP)
  0x002a 00042 (main.go:7)	MOVQ	"".slice+104(SP), AX   // *
  0x002f 00047 (main.go:7)	MOVQ	AX, 8(SP)
  0x0034 00052 (main.go:7)	MOVQ	"".slice+112(SP), AX   // *
  0x0039 00057 (main.go:7)	MOVQ	AX, 16(SP)
  0x003e 00062 (main.go:7)	PCDATA	$1, $1
  0x003e 00062 (main.go:7)	NOP
  0x0040 00064 (main.go:7)	CALL	runtime.convTslice(SB) // *
  0x0045 00069 (main.go:7)	MOVQ	24(SP), AX             // *
  0x004a 00074 (main.go:7)	XORPS	X0, X0
  0x004d 00077 (main.go:7)	MOVUPS	X0, ""..autotmp_13+64(SP)
  0x0052 00082 (main.go:7)	LEAQ	type.[]int(SB), CX
  0x0059 00089 (main.go:7)	MOVQ	CX, ""..autotmp_13+64(SP)
  0x005e 00094 (main.go:7)	MOVQ	AX, ""..autotmp_13+72(SP)
*/
