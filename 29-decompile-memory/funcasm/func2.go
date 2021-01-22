package main

func add2(a, b int) int {
	times := 2
	println(times)
	return (a - b) * times
}

func main() {
	sum := add2(1, 2)
	println(sum)
}

/*
"".add2 STEXT size=103 args=0x18 locals=0x10
  0x0000 00000 (func2.go:3)	TEXT	"".add2(SB), ABIInternal, $16-24
  0x0000 00000 (func2.go:3)	MOVQ	(TLS), CX   // *
  0x0009 00009 (func2.go:3)	CMPQ	SP, 16(CX)  // *
  0x000d 00013 (func2.go:3)	PCDATA	$0, $-2
  0x000d 00013 (func2.go:3)	JLS	96          // *
  0x000f 00015 (func2.go:3)	PCDATA	$0, $-1
  0x000f 00015 (func2.go:3)	SUBQ	$16, SP     // # 栈空间
  0x0013 00019 (func2.go:3)	MOVQ	BP, 8(SP)   // # 老的基址(保存 BP 寄存器的值到内存地址 8(SP))
  0x0018 00024 (func2.go:3)	LEAQ	8(SP), BP   // # 新的基址(拷贝 8(SP) 的内存引用地址值到 BP)
  ...
  0x001d 00029 (func2.go:5)	NOP
  0x0020 00032 (func2.go:5)	CALL	runtime.printlock(SB)
  0x0025 00037 (func2.go:5)	MOVQ	$2, (SP)
  0x002d 00045 (func2.go:5)	CALL	runtime.printint(SB)
  0x0032 00050 (func2.go:5)	CALL	runtime.printnl(SB)
  0x0037 00055 (func2.go:5)	CALL	runtime.printunlock(SB)
  0x003c 00060 (func2.go:6)	MOVQ	"".a+24(SP), AX
  0x0041 00065 (func2.go:6)	MOVQ	"".b+32(SP), CX
  0x0046 00070 (func2.go:6)	SUBQ	CX, AX
  0x0049 00073 (func2.go:6)	SHLQ	$1, AX
  0x004c 00076 (func2.go:6)	MOVQ	AX, "".~r2+40(SP)
  0x0051 00081 (func2.go:6)	MOVQ	8(SP), BP
  0x0056 00086 (func2.go:6)	ADDQ	$16, SP
  0x005a 00090 (func2.go:6)	RET
  0x005b 00091 (func2.go:6)	NOP
  0x005b 00091 (func2.go:3)	PCDATA	$1, $-1
  0x005b 00091 (func2.go:3)	PCDATA	$0, $-2
  0x005b 00091 (func2.go:3)	NOP
  0x0060 00096 (func2.go:3)	CALL	runtime.morestack_noctxt(SB) // *
  0x0065 00101 (func2.go:3)	PCDATA	$0, $-1                      // *
  0x0065 00101 (func2.go:3)	JMP	0                            // *
*/
