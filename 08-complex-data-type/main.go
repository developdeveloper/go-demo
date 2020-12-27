package main

import (
	"08-complex-data-type/keyvalue"
	"fmt"
	"unsafe"
)

func sliceSize() {
	s1 := []int{}
	fmt.Println(unsafe.Sizeof(s1)) // 24
}

func createSliceFromArray() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	s1 := arr1[:]
	fmt.Printf("%T %T\n", arr1, s1) // [5]int []int

	s2 := arr1[1:2] // 索引从1开始到2
	fmt.Println(s2) // [2]

	s3 := arr1[2:3]               // 索引从2开始到3
	fmt.Println(s3)               // [3]
	fmt.Println(len(s3), cap(s3)) // 1 3

	s4 := arr1[2:3:4]
	fmt.Println(s4)               // [3]
	fmt.Println(len(s4), cap(s4)) // 1 2

	s5 := arr1[2:3:5] // ok
	fmt.Println(s5)
	fmt.Println(len(s5), cap(s5)) // 1 3

	// s6 := arr1[2:3:6] // invalid slice index 6 (out of bounds for 5-element array)
	// fmt.Println(s6)

	s5 = append(s5, 4)
	fmt.Println(len(s5), cap(s5)) // 2 3
	s5 = append(s5, 5, 6, 7)      // 超过了原来的容量，会自动扩容
	fmt.Println(s5)               // [3 4 5 6 7]
	fmt.Println(len(s5), cap(s5)) // 5 6

	s7 := []int{2, 3, 4}
	fmt.Println(len(s7), cap(s7)) // 3 3
	copy(s7, []int{1, 2, 3, 4, 5})
	fmt.Println(s7)               // [1 2 3] 超出的元素没有，不会自动扩容
	fmt.Println(len(s7), cap(s7)) // 3 3

	s8 := []int{4: 5}
	fmt.Println(s8)               // [0 0 0 0 5]
	fmt.Println(len(s8), cap(s8)) // 5 5

	const size = 5
	const max = 10
	s9 := make([]int, size, max)
	fmt.Println(s9)               // [0 0 0 0 0]
	fmt.Println(len(s9), cap(s9)) // 5 10

	s10 := new([]int)
	*s10 = append(*s10, 1, 2, 3)
	fmt.Println(*s10)                 // [1 2 3]
	fmt.Println(len(*s10), cap(*s10)) // 3 4
}

func mapOperate() {
	myDict := map[string]string{}
	myDict["one"] = "1"
	myDict["two"] = "2"
	myDict["three"] = "3"

	fmt.Println(myDict)        // map[one:1 three:3 two:2]
	fmt.Println(myDict["one"]) // 1

	delete(myDict, "one")
	str, ok := myDict["one"]
	fmt.Println(ok)  // false
	fmt.Println(str) // 空

	fmt.Println(len(myDict)) // 2
	myDict["one"] = "1"
	for key, value := range myDict {
		fmt.Println(key, value)
	}

	for key := range myDict {
		fmt.Println(key, myDict[key])
	}
}

func emptyStructKey() {
	m := map[string]struct{}{}
	m["one"] = struct{}{}
}

func useDictionary() {
	var dict keyvalue.Dictionary
	fmt.Printf("%T\n", dict) // keyvalue.Dictionary
}

func main() {
	sliceSize()
	createSliceFromArray()

	mapOperate()
	useDictionary()
}
