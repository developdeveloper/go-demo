package array

import "fmt"

func initialSpecial() [5]int {
	arr := [5]int{1: 2, 2: 3}
	fmt.Println(len(arr))   // 输出 5
	fmt.Printf("%T\n", arr) // 输出 [5]int
	return arr
}

func updateFailed(arr [3]int) {
	fmt.Printf("%p\n", &arr) // 0xc000014340
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	fmt.Println(arr) // [1 1 1]
}

func updateSuccessed(ptrArr *[3]int) {
	fmt.Printf("%p\n", ptrArr) // 0xc0000cc140
	ptrArr[0] = 1
	ptrArr[1] = 1
	ptrArr[2] = 1
	fmt.Println(*ptrArr) // [1 1 1]
}

func printHugeArray1(arr [10000000]int) {
	// 啥也不干
}

func printHugeArray2(arr *[10000000]int) {
	// 啥也不干
}

func printHugeSlice(arr []int) {
	// 啥也不干
}
