package tailrecfunc

func factor(num uint64) uint64 {
	if num == 1 {
		return 1
	}

	return num * factor(num-1)
}

/*
"".factor STEXT size=108 args=0x10 locals=0x18
	0x0000 00000 (factor.go:3)	TEXT	"".factor(SB), ABIInternal, $24-16
	0x0000 00000 (factor.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (factor.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (factor.go:3)	PCDATA	$0, $-2
	0x000d 00013 (factor.go:3)	JLS	101
	0x000f 00015 (factor.go:3)	PCDATA	$0, $-1
	0x000f 00015 (factor.go:3)	SUBQ	$24, SP
	0x0013 00019 (factor.go:3)	MOVQ	BP, 16(SP)
	0x0018 00024 (factor.go:3)	LEAQ	16(SP), BP
	0x001d 00029 (factor.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (factor.go:4)	MOVQ	"".num+32(SP), AX
	0x0022 00034 (factor.go:4)	CMPQ	AX, $1
	0x0026 00038 (factor.go:4)	JNE	59
	0x0028 00040 (factor.go:5)	MOVQ	$1, "".~r1+40(SP)
	0x0031 00049 (factor.go:5)	MOVQ	16(SP), BP
	0x0036 00054 (factor.go:5)	ADDQ	$24, SP
	0x003a 00058 (factor.go:5)	RET
	0x003b 00059 (factor.go:8)	LEAQ	-1(AX), CX
	0x003f 00063 (factor.go:8)	MOVQ	CX, (SP)
	0x0043 00067 (factor.go:8)	PCDATA	$1, $0
	0x0043 00067 (factor.go:8)	CALL	"".factor(SB)
	0x0048 00072 (factor.go:8)	MOVQ	8(SP), AX
	0x004d 00077 (factor.go:8)	MOVQ	"".num+32(SP), CX
	0x0052 00082 (factor.go:8)	IMULQ	AX, CX
	0x0056 00086 (factor.go:8)	MOVQ	CX, "".~r1+40(SP)
	0x005b 00091 (factor.go:8)	MOVQ	16(SP), BP
	0x0060 00096 (factor.go:8)	ADDQ	$24, SP
	0x0064 00100 (factor.go:8)	RET
	0x0065 00101 (factor.go:8)	NOP
	0x0065 00101 (factor.go:3)	PCDATA	$1, $-1
	0x0065 00101 (factor.go:3)	PCDATA	$0, $-2
	0x0065 00101 (factor.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x006a 00106 (factor.go:3)	PCDATA	$0, $-1
	0x006a 00106 (factor.go:3)	JMP	0
*/
