package entity

import (
	"fmt"
	"reflect"
	"unsafe"
)

func lookupEmptyStruct() {
	a := struct{}{}
	b := struct{}{}

	if reflect.DeepEqual(a, b) {
		fmt.Printf("%p\n", &a) // 0x125a7d0
		fmt.Printf("%p\n", &b) // 0x125a7d0
		fmt.Println("空结构体都相等")
		fmt.Printf("%d\n", unsafe.Sizeof(a)) // 0
		fmt.Printf("%d\n", unsafe.Sizeof(b)) // 0
	}
}