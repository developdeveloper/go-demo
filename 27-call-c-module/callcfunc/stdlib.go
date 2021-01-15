package main

import "C"

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("protohello, cgo"))
}
