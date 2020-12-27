package array

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_initialSpecial(t *testing.T) {
	arr := initialSpecial()
	// 数组不是简单类型，使用 reflect.DeepEqual 来判断值是否相等
	if !reflect.DeepEqual(arr, [5]int{0, 2, 3, 0, 0}) {
		t.Fatal("数组初始化错误")
	}
}

func Test_updateFailed(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr) // 0xc000014320
	updateFailed(arr)
	fmt.Println(arr) // [1 2 3]
	if !reflect.DeepEqual(arr, [3]int{1, 2, 3}) {
		t.Fatal("函数改变了原数组")
	}
}

func Test_updateSuccessed(t *testing.T) {
	arr := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr) // 0xc0000cc140
	updateSuccessed(&arr)
	fmt.Println(arr) // [1 1 1]
	if !reflect.DeepEqual(arr, [3]int{1, 1, 1}) {
		t.Fatal("函数改变了原数组")
	}
}

func Benchmark_printHugeArray1(b *testing.B) {
	arr := [10000000]int{}
	for i := 0; i < b.N; i++ {
		printHugeArray1(arr)
	}
}

func Benchmark_printHugeArray2(b *testing.B) {
	arr := [10000000]int{}
	for i := 0; i < b.N; i++ {
		printHugeArray2(&arr)
	}
}

func Benchmark_printHugeSlice(b *testing.B) {
	arr := [10000000]int{}
	for i := 0; i < b.N; i++ {
		printHugeSlice(arr[:])
	}
}
