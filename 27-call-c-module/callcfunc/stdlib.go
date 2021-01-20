package main

//#include <stdio.h>
//#include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	cstr := C.CString("hi, cgo lib")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))

	str := C.GoString(cstr)
	fmt.Println(str)
}
