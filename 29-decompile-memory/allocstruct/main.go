package main

type Person struct{}

func main() {
	// var p *Person = &Person{}
	p := &Person{}
	// p := new(Person)
	println(p)
}

/*
// p := &Person{}
0x0020 00032 (main.go:8)	CALL	runtime.printlock(SB)
0x0025 00037 (main.go:8)	LEAQ	""..autotmp_2+8(SP), AX
0x002a 00042 (main.go:8)	MOVQ	AX, (SP)
0x002e 00046 (main.go:8)	CALL	runtime.printpointer(SB)
0x0033 00051 (main.go:8)	CALL	runtime.printnl(SB)
0x0038 00056 (main.go:8)	CALL	runtime.printunlock(SB)
0x003d 00061 (main.go:9)	MOVQ	8(SP), BP
0x0042 00066 (main.go:9)	ADDQ	$16, SP
0x0046 00070 (main.go:9)	RET

// p := new(Person)
0x0020 00032 (main.go:9)	CALL	runtime.printlock(SB)
0x0025 00037 (main.go:9)	LEAQ	""..autotmp_1+8(SP), AX
0x002a 00042 (main.go:9)	MOVQ	AX, (SP)
0x002e 00046 (main.go:9)	CALL	runtime.printpointer(SB)
0x0033 00051 (main.go:9)	CALL	runtime.printnl(SB)
0x0038 00056 (main.go:9)	CALL	runtime.printunlock(SB)
0x003d 00061 (main.go:10)	MOVQ	8(SP), BP
0x0042 00066 (main.go:10)	ADDQ	$16, SP
0x0046 00070 (main.go:10)	RET

0x0020 00032 (main.go:5)	CALL	runtime.printlock(SB)
0x0025 00037 (main.go:5)	LEAQ	""..autotmp_2+8(SP), AX
0x002a 00042 (main.go:5)	MOVQ	AX, (SP)
0x002e 00046 (main.go:5)	CALL	runtime.printpointer(SB)
0x0033 00051 (main.go:5)	CALL	runtime.printnl(SB)
0x0038 00056 (main.go:5)	CALL	runtime.printunlock(SB)
0x003d 00061 (main.go:6)	MOVQ	8(SP), BP
0x0042 00066 (main.go:6)	ADDQ	$16, SP
0x0046 00070 (main.go:6)	RET
*/
