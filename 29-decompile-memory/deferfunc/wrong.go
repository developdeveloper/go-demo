package main

func f() {
	i := 0

	defer func() {
		println(i) // 1
	}()

	i++
}

func main() {
	f()
}

/*
	0x0025 00037 (wrong.go:3)	MOVB	$0, ""..autotmp_2+15(SP)
	0x002a 00042 (wrong.go:4)	MOVQ	$0, "".i+16(SP)
	0x0033 00051 (wrong.go:6)	LEAQ	"".f.func1·f(SB), AX
	0x003a 00058 (wrong.go:6)	MOVQ	AX, ""..autotmp_3+32(SP)
	0x003f 00063 (wrong.go:6)	LEAQ	"".i+16(SP), AX
	0x0044 00068 (wrong.go:6)	MOVQ	AX, ""..autotmp_4+24(SP)
	0x0049 00073 (wrong.go:6)	MOVB	$1, ""..autotmp_2+15(SP)
	0x004e 00078 (wrong.go:10)	MOVQ	"".i+16(SP), AX
	0x0053 00083 (wrong.go:10)	INCQ	AX
	0x0056 00086 (wrong.go:10)	MOVQ	AX, "".i+16(SP)
	0x005b 00091 (wrong.go:11)	MOVB	$0, ""..autotmp_2+15(SP)
	0x0060 00096 (wrong.go:11)	MOVQ	""..autotmp_4+24(SP), AX
	0x0065 00101 (wrong.go:11)	MOVQ	AX, (SP)
	0x0069 00105 (wrong.go:11)	PCDATA	$1, $1
	0x0069 00105 (wrong.go:11)	CALL	"".f.func1(SB)
	0x006e 00110 (wrong.go:11)	MOVQ	40(SP), BP
	0x0073 00115 (wrong.go:11)	ADDQ	$48, SP
	0x0077 00119 (wrong.go:11)	RET
	0x0078 00120 (wrong.go:11)	CALL	runtime.deferreturn(SB)
	0x007d 00125 (wrong.go:11)	MOVQ	40(SP), BP
	0x0082 00130 (wrong.go:11)	ADDQ	$48, SP
	0x0086 00134 (wrong.go:11)	RET
*/

/*
"".f.func1 STEXT size=86 args=0x8 locals=0x10
	0x0000 00000 (wrong.go:6)	TEXT	"".f.func1(SB), ABIInternal, $16-8
	0x0000 00000 (wrong.go:6)	MOVQ	(TLS), CX
	0x0009 00009 (wrong.go:6)	CMPQ	SP, 16(CX)
	0x000d 00013 (wrong.go:6)	PCDATA	$0, $-2
	0x000d 00013 (wrong.go:6)	JLS	79
	0x000f 00015 (wrong.go:6)	PCDATA	$0, $-1
	0x000f 00015 (wrong.go:6)	SUBQ	$16, SP
	0x0013 00019 (wrong.go:6)	MOVQ	BP, 8(SP)
	0x0018 00024 (wrong.go:6)	LEAQ	8(SP), BP
	0x001d 00029 (wrong.go:6)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (wrong.go:6)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (wrong.go:7)	PCDATA	$1, $0
	0x001d 00029 (wrong.go:7)	NOP
	0x0020 00032 (wrong.go:7)	CALL	runtime.printlock(SB)
	0x0025 00037 (wrong.go:7)	MOVQ	"".&i+24(SP), AX
	0x002a 00042 (wrong.go:7)	MOVQ	(AX), AX
	0x002d 00045 (wrong.go:7)	MOVQ	AX, (SP)
	0x0031 00049 (wrong.go:7)	PCDATA	$1, $1
	0x0031 00049 (wrong.go:7)	CALL	runtime.printint(SB)
	0x0036 00054 (wrong.go:7)	CALL	runtime.printnl(SB)
	0x003b 00059 (wrong.go:7)	NOP
	0x0040 00064 (wrong.go:7)	CALL	runtime.printunlock(SB)
	0x0045 00069 (wrong.go:8)	MOVQ	8(SP), BP
	0x004a 00074 (wrong.go:8)	ADDQ	$16, SP
	0x004e 00078 (wrong.go:8)	RET
*/
