package alglib

import (
	"fmt"
	"sort"
)

// SortNumberTest 排序
func SortNumberTest() {
	numbers := []int{4, 5, 2, 5, 6, 7, 8, 0}
	sort.Ints(numbers)
	fmt.Println(numbers)

	// 要求 asc order 顺序，参数是 sorted slice 已经排好序
	fmt.Println(sort.SearchInts(numbers, 2))

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)
}
