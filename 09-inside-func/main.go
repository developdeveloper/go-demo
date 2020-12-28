package main

import (
	_ "09-inside-func/initial"
	"fmt"
)

func init() {
	fmt.Println("init called in main package")
}

func main() {
	fmt.Println("main called")

	// initial.DoNothing()
}
