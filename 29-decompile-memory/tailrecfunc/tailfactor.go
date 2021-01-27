package tailrecfunc

func tailFactor(num, current uint64) uint64 {
	if num == 1 {
		return current
	}

	return tailFactor(num-1, num*current)
}

/*
"".tailFactor STEXT size=114 args=0x18 locals=0x20
	0x0000 00000 (tailfactor.go:3)	TEXT	"".tailFactor(SB), ABIInternal, $32-24
	0x0000 00000 (tailfactor.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (tailfactor.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (tailfactor.go:3)	PCDATA	$0, $-2
	0x000d 00013 (tailfactor.go:3)	JLS	107
	0x000f 00015 (tailfactor.go:3)	PCDATA	$0, $-1
	0x000f 00015 (tailfactor.go:3)	SUBQ	$32, SP
	0x0013 00019 (tailfactor.go:3)	MOVQ	BP, 24(SP)
	0x0018 00024 (tailfactor.go:3)	LEAQ	24(SP), BP
	0x001d 00029 (tailfactor.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tailfactor.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (tailfactor.go:4)	MOVQ	"".num+40(SP), AX
	0x0022 00034 (tailfactor.go:4)	CMPQ	AX, $1
	0x0026 00038 (tailfactor.go:4)	JNE	60
	0x0028 00040 (tailfactor.go:5)	MOVQ	"".current+48(SP), AX
	0x002d 00045 (tailfactor.go:5)	MOVQ	AX, "".~r2+56(SP)
	0x0032 00050 (tailfactor.go:5)	MOVQ	24(SP), BP
	0x0037 00055 (tailfactor.go:5)	ADDQ	$32, SP
	0x003b 00059 (tailfactor.go:5)	RET
	0x003c 00060 (tailfactor.go:8)	LEAQ	-1(AX), CX
	0x0040 00064 (tailfactor.go:8)	MOVQ	CX, (SP)
	0x0044 00068 (tailfactor.go:8)	MOVQ	"".current+48(SP), CX
	0x0049 00073 (tailfactor.go:8)	IMULQ	CX, AX
	0x004d 00077 (tailfactor.go:8)	MOVQ	AX, 8(SP)
	0x0052 00082 (tailfactor.go:8)	PCDATA	$1, $0
	0x0052 00082 (tailfactor.go:8)	CALL	"".tailFactor(SB)
	0x0057 00087 (tailfactor.go:8)	MOVQ	16(SP), AX
	0x005c 00092 (tailfactor.go:8)	MOVQ	AX, "".~r2+56(SP)
	0x0061 00097 (tailfactor.go:8)	MOVQ	24(SP), BP
	0x0066 00102 (tailfactor.go:8)	ADDQ	$32, SP
	0x006a 00106 (tailfactor.go:8)	RET
	0x006b 00107 (tailfactor.go:8)	NOP
	0x006b 00107 (tailfactor.go:3)	PCDATA	$1, $-1
	0x006b 00107 (tailfactor.go:3)	PCDATA	$0, $-2
	0x006b 00107 (tailfactor.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0070 00112 (tailfactor.go:3)	PCDATA	$0, $-1
	0x0070 00112 (tailfactor.go:3)	JMP	0
*/
