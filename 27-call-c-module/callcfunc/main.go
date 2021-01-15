package main

/*
#include <stdio.h>

static void Say(const char* s) {
	puts(s);
}
*/
import "C"

func main() {
	C.Say(C.CString("hi cgo"))
}
