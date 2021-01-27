package main

//go:noinline
func printArray(data [3]int) {
	print(data[0])
}

func main() {
	printArray([3]int{1, 2, 3})
}

/*
"".printArray STEXT size=80 args=0x18 locals=0x18
	0x0000 00000 (array.go:3)	TEXT	"".printArray(SB), ABIInternal, $24-24
	0x0000 00000 (array.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (array.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (array.go:3)	PCDATA	$0, $-2
	0x000d 00013 (array.go:3)	JLS	73
	0x000f 00015 (array.go:3)	PCDATA	$0, $-1
	0x000f 00015 (array.go:3)	SUBQ	$24, SP
	0x0013 00019 (array.go:3)	MOVQ	BP, 16(SP)
	0x0018 00024 (array.go:3)	LEAQ	16(SP), BP
	0x001d 00029 (array.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (array.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (array.go:4)	MOVQ	"".data+32(SP), AX
	0x0022 00034 (array.go:4)	MOVQ	AX, ""..autotmp_2+8(SP)
	0x0027 00039 (array.go:4)	PCDATA	$1, $0
	0x0027 00039 (array.go:4)	CALL	runtime.printlock(SB)
	0x002c 00044 (array.go:4)	MOVQ	""..autotmp_2+8(SP), AX
	0x0031 00049 (array.go:4)	MOVQ	AX, (SP)
	0x0035 00053 (array.go:4)	CALL	runtime.printint(SB)
	0x003a 00058 (array.go:4)	CALL	runtime.printunlock(SB)
	0x003f 00063 (array.go:5)	MOVQ	16(SP), BP
	0x0044 00068 (array.go:5)	ADDQ	$24, SP
	0x0048 00072 (array.go:5)	RET
	0x0049 00073 (array.go:5)	NOP
	0x0049 00073 (array.go:3)	PCDATA	$1, $-1
	0x0049 00073 (array.go:3)	PCDATA	$0, $-2
	0x0049 00073 (array.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x004e 00078 (array.go:3)	PCDATA	$0, $-1
	0x004e 00078 (array.go:3)	JMP	0

"".main STEXT size=77 args=0x0 locals=0x20
	0x0000 00000 (array.go:8)	TEXT	"".main(SB), ABIInternal, $32-0
	0x0000 00000 (array.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (array.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (array.go:8)	PCDATA	$0, $-2
	0x000d 00013 (array.go:8)	JLS	70
	0x000f 00015 (array.go:8)	PCDATA	$0, $-1
	0x000f 00015 (array.go:8)	SUBQ	$32, SP
	0x0013 00019 (array.go:8)	MOVQ	BP, 24(SP)
	0x0018 00024 (array.go:8)	LEAQ	24(SP), BP
	0x001d 00029 (array.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (array.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (array.go:9)	MOVQ	$1, (SP)
	0x0025 00037 (array.go:9)	MOVQ	$2, 8(SP)
	0x002e 00046 (array.go:9)	MOVQ	$3, 16(SP)
	0x0037 00055 (array.go:9)	PCDATA	$1, $0
	0x0037 00055 (array.go:9)	CALL	"".printArray(SB)
	0x003c 00060 (array.go:10)	MOVQ	24(SP), BP
	0x0041 00065 (array.go:10)	ADDQ	$32, SP
	0x0045 00069 (array.go:10)	RET
	0x0046 00070 (array.go:10)	NOP
	0x0046 00070 (array.go:8)	PCDATA	$1, $-1
	0x0046 00070 (array.go:8)	PCDATA	$0, $-2
	0x0046 00070 (array.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x004b 00075 (array.go:8)	PCDATA	$0, $-1
	0x004b 00075 (array.go:8)	JMP	0
*/
