// say.go

package main

import "C"

import "fmt"

//export Say
func Say(s *C.char) {
	fmt.Print(C.GoString(s))
}
