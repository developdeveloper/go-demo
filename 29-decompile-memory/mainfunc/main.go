package main

var str = "hello"

func main() {
	print(str)
}

/*
"".main STEXT size=86 args=0x0 locals=0x18
	0x0000 00000 (main.go:7)	TEXT	"".main(SB), ABIInternal, $24-0
	0x0000 00000 (main.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:7)	PCDATA	$0, $-2
	0x000d 00013 (main.go:7)	JLS	79
	0x000f 00015 (main.go:7)	PCDATA	$0, $-1
	0x000f 00015 (main.go:7)	SUBQ	$24, SP
	0x0013 00019 (main.go:7)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:7)	LEAQ	16(SP), BP
	0x001d 00029 (main.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:8)	PCDATA	$1, $0
	0x001d 00029 (main.go:8)	NOP
	0x0020 00032 (main.go:8)	CALL	runtime.printlock(SB)
	0x0025 00037 (main.go:8)	LEAQ	go.string."darwin\n"(SB), AX
	0x002c 00044 (main.go:8)	MOVQ	AX, (SP)
	0x0030 00048 (main.go:8)	MOVQ	$7, 8(SP)
	0x0039 00057 (main.go:8)	CALL	runtime.printstring(SB)
	0x003e 00062 (main.go:8)	NOP
	0x0040 00064 (main.go:8)	CALL	runtime.printunlock(SB)
	0x0045 00069 (main.go:9)	MOVQ	16(SP), BP
	0x004a 00074 (main.go:9)	ADDQ	$24, SP
	0x004e 00078 (main.go:9)	RET
	0x004f 00079 (main.go:9)	NOP
	0x004f 00079 (main.go:7)	PCDATA	$1, $-1
	0x004f 00079 (main.go:7)	PCDATA	$0, $-2
	0x004f 00079 (main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0054 00084 (main.go:7)	PCDATA	$0, $-1
	0x0054 00084 (main.go:7)	JMP	0
*/
