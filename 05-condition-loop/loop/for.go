package loop

import "fmt"

// FullFor 完整的 for 定义
func FullFor() {
	sum := 0

	// 初始值; 循环条件; 增量/减量
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// ConditionFor 只有循环条件
func ConditionFor() {
	sum := 1
	// for ; sum < 10 ; {
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}

// DeadFor 死循环
func DeadFor() {
	for true {
		// 无限循环
	}

	for {
		// 无限循环
	}
}

//BreakContinueFor 跳过和中断
func BreakContinueFor() {
	for i := 0; i < 100; i++ {
		// 余数是 0 是偶数，跳过
		if i%2 == 0 {
			continue
		}

		// 大于 10 就中断循环
		if i > 10 {
			break
		}

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// ForRange 安全的迭代循环
func ForRange() {
	str := "this is a test"
	for index, ch := range str {
		// 跳过空格不处理
		if ch != 0x20 {
			fmt.Printf("%d%c ", index, ch)
		}
	}
	fmt.Println()
}

// ForRange2 安全的迭代循环
func ForRange2() {
	str := "this is a test"
	for index := range str {
		fmt.Printf("%d ", index)
	}
	fmt.Println()
}

// ForRange3 安全的迭代循环
func ForRange3() {
	str := "this is a test"
	for _, ch := range str {
		fmt.Printf("%c", ch)
	}
	fmt.Println()
}
