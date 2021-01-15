package main

//#include "say.h"
import "C"

func main() {
	C.Say(C.CString("hi, cgo"))
}
