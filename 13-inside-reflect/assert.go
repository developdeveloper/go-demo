package main

import "fmt"

func safeAssert() {
	var i interface{}
	//_ = i.(interface{})
	// _ = i.(string) // panic: interface conversion: interface {} is nil, not string
	str, ok := i.(string)
	if ok {
		fmt.Println(str)
	}
}
