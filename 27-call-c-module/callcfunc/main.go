package main

/*
#include <stdio.h>

void Say(const char* s) {
	puts(s);
}
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("hi, cgo")
	C.Say(cstr)
	C.free(unsafe.Pointer(cstr))
}
