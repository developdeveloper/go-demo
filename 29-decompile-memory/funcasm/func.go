package main

func add(a, b int) int {
	return a + b
}

func main() {
	sum := add(1, 2)
	println(sum)
}

/*
"".add STEXT nosplit size=19 args=0x18 locals=0x0
  0x0000 00000 (func.go:3)	TEXT	"".add(SB), NOSPLIT|ABIInternal, $0-24
  ...
  0x0000 00000 (func.go:4)	MOVQ	"".b+16(SP), AX
  0x0005 00005 (func.go:4)	MOVQ	"".a+8(SP), CX
  0x000a 00010 (func.go:4)	ADDQ	CX, AX
  0x000d 00013 (func.go:4)	MOVQ	AX, "".~r2+24(SP)
  0x0012 00018 (func.go:4)	RET
*/
