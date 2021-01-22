package main

//go:noinline
func swap(a, b int) (int, int) {
	return b + 1, a + 1
}

func main() {
	_, _ = swap(0, 1)
}

/*
"".swap STEXT nosplit size=27 args=0x20 locals=0x0
  0x0000 00000 (main.go:4)	TEXT	"".swap(SB), NOSPLIT|ABIInternal, $0-32
  0x0000 00000 (main.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x0000 00000 (main.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x0000 00000 (main.go:5)	MOVQ	"".b+16(SP), AX
  0x0005 00005 (main.go:5)	INCQ	AX
  0x0008 00008 (main.go:5)	MOVQ	AX, "".~r2+24(SP) // *
  0x000d 00013 (main.go:5)	MOVQ	"".a+8(SP), AX
  0x0012 00018 (main.go:5)	INCQ	AX
  0x0015 00021 (main.go:5)	MOVQ	AX, "".~r3+32(SP) // *
	0x001a 00026 (main.go:5)	RET

"".main STEXT size=71 args=0x0 locals=0x28
  0x0000 00000 (main.go:8)	TEXT	"".main(SB), ABIInternal, $40-0
  0x0000 00000 (main.go:8)	MOVQ	(TLS), CX
  0x0009 00009 (main.go:8)	CMPQ	SP, 16(CX)
  0x000d 00013 (main.go:8)	PCDATA	$0, $-2
  0x000d 00013 (main.go:8)	JLS	64
  0x000f 00015 (main.go:8)	PCDATA	$0, $-1
  0x000f 00015 (main.go:8)	SUBQ	$40, SP
  0x0013 00019 (main.go:8)	MOVQ	BP, 32(SP) // 基址
  0x0018 00024 (main.go:8)	LEAQ	32(SP), BP
  0x001d 00029 (main.go:8)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x001d 00029 (main.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x001d 00029 (main.go:9)	MOVQ	$0, (SP)   // *
  0x0025 00037 (main.go:9)	MOVQ	$1, 8(SP)  // *
  0x002e 00046 (main.go:9)	PCDATA	$1, $0
  0x002e 00046 (main.go:9)	CALL	"".swap(SB)
  0x0033 00051 (main.go:10)	MOVQ	32(SP), BP // *
  0x0038 00056 (main.go:10)	ADDQ	$40, SP    // *
  0x003c 00060 (main.go:10)	RET
  0x003d 00061 (main.go:10)	NOP
  0x003d 00061 (main.go:8)	PCDATA	$1, $-1
  0x003d 00061 (main.go:8)	PCDATA	$0, $-2
  0x003d 00061 (main.go:8)	NOP
  0x0040 00064 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
  0x0045 00069 (main.go:8)	PCDATA	$0, $-1
  0x0045 00069 (main.go:8)	JMP	0
*/
