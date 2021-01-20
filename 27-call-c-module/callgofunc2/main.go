package main

//void Say(char* s);
import "C"

import (
	"fmt"
)

func main() {
	C.Say(C.CString("hi, cgo"))
}

//export Say
func Say(s *C.char) {
	fmt.Println(C.GoString(s))
}
