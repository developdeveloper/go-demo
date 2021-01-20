package main

//void Say(_GoString_ s);
import "C"

import (
	"fmt"
)

func main() {
	C.Say("hi, cgo")
}

//export Say
func Say(s string) {
	fmt.Print(s)
}
