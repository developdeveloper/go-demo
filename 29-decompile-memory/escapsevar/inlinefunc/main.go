package main

type point struct {
	x, y int
}

/*
	0x0000 00000 (main.go:8)	MOVQ	"".p+8(SP), AX
	0x0005 00005 (main.go:8)	MOVQ	$10, (AX)
	0x000c 00012 (main.go:9)	MOVQ	$10, 8(AX)
	0x0014 00020 (main.go:10)	RET
*/
func move(p *point) {
	p.x = 10
	p.y = 10
}

/*
	0x001d 00029 (main.go:15)	XORPS	X0, X0
	0x0020 00032 (main.go:14)	MOVUPS	X0, ""..autotmp_2+8(SP)
	0x0025 00037 (<unknown line number>)	NOP
	0x0025 00037 (main.go:8)	MOVQ	$10, ""..autotmp_2+8(SP)
	0x002e 00046 (main.go:9)	MOVQ	$10, ""..autotmp_2+16(SP)
	0x0037 00055 (main.go:16)	PCDATA	$1, $0
	0x0037 00055 (main.go:16)	CALL	runtime.printlock(SB)
	0x003c 00060 (main.go:16)	LEAQ	""..autotmp_2+8(SP), AX
	0x0041 00065 (main.go:16)	MOVQ	AX, (SP)
	0x0045 00069 (main.go:16)	CALL	runtime.printpointer(SB)
	0x004a 00074 (main.go:16)	CALL	runtime.printunlock(SB)
	0x004f 00079 (main.go:17)	MOVQ	24(SP), BP
	0x0054 00084 (main.go:17)	ADDQ	$32, SP
	0x0058 00088 (main.go:17)	RET
*/

/*
# go build -gcflags=-m main.go
./main.go:13:6: can inline move
./main.go:34:6: can inline main
./main.go:37:6: inlining call to move
./main.go:13:11: p does not escape
./main.go:36:10: new(point) does not escape
*/
func main() {
	// 分配到栈上，注意 move 函数被内敛了
	p := new(point)
	move(p)
	print(p)
}
