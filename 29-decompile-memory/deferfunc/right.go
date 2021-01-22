package main

//go:noinline
func f() int {
	i := 0

	defer func(val int) {
		val++
		println(val)
	}(i)

	return i
}

func main() {
	println(f())
}

/*
	0x0029 00041 (defer.go:3)	MOVB	$0, ""..autotmp_3+15(SP)
	0x002e 00046 (defer.go:3)	MOVQ	$0, "".~r0+48(SP)
	0x0037 00055 (defer.go:6)	LEAQ	"".f.func1·f(SB), AX
	0x003e 00062 (defer.go:6)	MOVQ	AX, ""..autotmp_4+24(SP)
	0x0043 00067 (defer.go:6)	MOVQ	$0, ""..autotmp_5+16(SP)
	0x004c 00076 (defer.go:6)	MOVB	$1, ""..autotmp_3+15(SP)
	0x0051 00081 (defer.go:10)	MOVQ	$0, "".~r0+48(SP)
	0x005a 00090 (defer.go:10)	MOVB	$0, ""..autotmp_3+15(SP)
	0x005f 00095 (defer.go:10)	MOVQ	""..autotmp_5+16(SP), AX
	0x0064 00100 (defer.go:10)	MOVQ	AX, (SP)
	0x0068 00104 (defer.go:10)	PCDATA	$1, $1
	0x0068 00104 (defer.go:10)	CALL	"".f.func1(SB)
	0x006d 00109 (defer.go:10)	MOVQ	32(SP), BP
	0x0072 00114 (defer.go:10)	ADDQ	$40, SP
	0x0076 00118 (defer.go:10)	RET
	0x0077 00119 (defer.go:10)	CALL	runtime.deferreturn(SB)
	0x007c 00124 (defer.go:10)	MOVQ	32(SP), BP
	0x0081 00129 (defer.go:10)	ADDQ	$40, SP
	0x0085 00133 (defer.go:10)	RET
	0x0086 00134 (defer.go:10)	NOP
	0x0086 00134 (defer.go:3)	PCDATA	$1, $-1
	0x0086 00134 (defer.go:3)	PCDATA	$0, $-2
	0x0086 00134 (defer.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x008b 00139 (defer.go:3)	PCDATA	$0, $-1
	0x008b 00139 (defer.go:3)	JMP	0
*/
