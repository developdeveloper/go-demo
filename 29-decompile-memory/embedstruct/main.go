package main

type Rank struct {
	Level int
}

type Address struct {
	Rank
	StreetNo int
}

//go:noinline
func printAddr(addr Address) {
	println(addr.Level)
	println(addr.StreetNo)
}

func main() {
	printAddr(Address{Rank: Rank{Level: 1}, StreetNo: 1})
}

/*
"".printAddr STEXT size=110 args=0x10 locals=0x10
	0x0000 00000 (main.go:13)	TEXT	"".printAddr(SB), ABIInternal, $16-16
	0x0000 00000 (main.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:13)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:13)	PCDATA	$0, $-2
	0x000d 00013 (main.go:13)	JLS	103
	0x000f 00015 (main.go:13)	PCDATA	$0, $-1
	0x000f 00015 (main.go:13)	SUBQ	$16, SP
	0x0013 00019 (main.go:13)	MOVQ	BP, 8(SP)
	0x0018 00024 (main.go:13)	LEAQ	8(SP), BP
	0x001d 00029 (main.go:13)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:13)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:14)	PCDATA	$1, $0
	0x001d 00029 (main.go:14)	NOP
	0x0020 00032 (main.go:14)	CALL	runtime.printlock(SB)
	0x0025 00037 (main.go:14)	MOVQ	"".addr+24(SP), AX
	0x002a 00042 (main.go:14)	MOVQ	AX, (SP)
	0x002e 00046 (main.go:14)	CALL	runtime.printint(SB)
	0x0033 00051 (main.go:14)	CALL	runtime.printnl(SB)
	0x0038 00056 (main.go:14)	CALL	runtime.printunlock(SB)
	0x003d 00061 (main.go:14)	NOP
	0x0040 00064 (main.go:15)	CALL	runtime.printlock(SB)
	0x0045 00069 (main.go:15)	MOVQ	"".addr+32(SP), AX
	0x004a 00074 (main.go:15)	MOVQ	AX, (SP)
	0x004e 00078 (main.go:15)	CALL	runtime.printint(SB)
	0x0053 00083 (main.go:15)	CALL	runtime.printnl(SB)
	0x0058 00088 (main.go:15)	CALL	runtime.printunlock(SB)
	0x005d 00093 (main.go:16)	MOVQ	8(SP), BP
	0x0062 00098 (main.go:16)	ADDQ	$16, SP
	0x0066 00102 (main.go:16)	RET
	0x0067 00103 (main.go:16)	NOP
	0x0067 00103 (main.go:13)	PCDATA	$1, $-1
	0x0067 00103 (main.go:13)	PCDATA	$0, $-2
	0x0067 00103 (main.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x006c 00108 (main.go:13)	PCDATA	$0, $-1
	0x006c 00108 (main.go:13)	JMP	0

"".main STEXT size=71 args=0x0 locals=0x18
	0x0000 00000 (main.go:18)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:18)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:18)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:18)	PCDATA	$0, $-2
	0x000d 00013 (main.go:18)	JLS	64
	0x000f 00015 (main.go:18)	PCDATA	$0, $-1
	0x000f 00015 (main.go:18)	SUBQ	$24, SP
	0x0013 00019 (main.go:18)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:18)	LEAQ	16(SP), BP
	0x001d 00029 (main.go:18)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:18)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:19)	MOVQ	$1, (SP)
	0x0025 00037 (main.go:19)	MOVQ	$1, 8(SP)
	0x002e 00046 (main.go:19)	PCDATA	$1, $0
	0x002e 00046 (main.go:19)	CALL	"".printAddr(SB)
	0x0033 00051 (main.go:20)	MOVQ	16(SP), BP
	0x0038 00056 (main.go:20)	ADDQ	$24, SP
	0x003c 00060 (main.go:20)	RET
	0x003d 00061 (main.go:20)	NOP
	0x003d 00061 (main.go:18)	PCDATA	$1, $-1
	0x003d 00061 (main.go:18)	PCDATA	$0, $-2
	0x003d 00061 (main.go:18)	NOP
	0x0040 00064 (main.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x0045 00069 (main.go:18)	PCDATA	$0, $-1
	0x0045 00069 (main.go:18)	JMP	0
*/
