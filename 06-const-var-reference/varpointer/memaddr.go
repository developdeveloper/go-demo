package varpointer

import "fmt"

// PrintAddr 打印变量的地址
func PrintAddr() {
	score := 10
	fmt.Printf("%p\n", &score)
}

// HoldVarAddr 持有变量的内存地址
func HoldVarAddr() {
	score := 10
	addr := &score
	*addr = 20
	fmt.Println(score) // 打印 20
}

// 不能修改入参 score 参数值
// 函数调用的时候参数被拷贝
func cannotUpdateScore(score int) {
	fmt.Println(&score) // 取参数的地址 0xc00001c138
	score++
	fmt.Println(score) // 11
}

// IncScoreFailed 调高分数
func IncScoreFailed() {
	score := 10
	fmt.Println(&score) // 0xc00001c130
	cannotUpdateScore(score)
	fmt.Println(score) // 10
}

// 可以修改入参 score 参数值
// 函数调用的时候参数被拷贝
func canUpdateScore(ptrScore *int) {
	fmt.Println(ptrScore) // 已经是地址了 0xc00001c140
	*ptrScore++
	fmt.Println(*ptrScore) // 11
}

// IncScoreSuccessed 调高分数
func IncScoreSuccessed() {
	score := 10
	fmt.Println(&score) // 0xc00001c140
	canUpdateScore(&score)
	fmt.Println(score) // 11
}
