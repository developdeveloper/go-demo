package branch

import "fmt"

// FullSwitch 挨个顺序匹配
func FullSwitch(rank string) {
	switch rank {
	case "gold":
		fmt.Println("1st")
	case "silver", "bronze":
		fmt.Println("2nd or 3rd")
	default:
		// 匹配不到才会执行
		fmt.Println("未获奖")
	}
}

// FallThrough 执行匹配到下一个 case 语句(无论它的条件是真是假)
func FallThrough() {
	switch {
	case true:
		fmt.Println("我是匹配到的")
		fallthrough
	case false:
		fmt.Println("虽然我是 false，但我会被打印")
		fallthrough
	case false:
		fmt.Println("虽然我也是 false，但我也会被打印")
		// fallthrough
		// 如果放开这个 fallthrough，下面的 default 也会被执行
	default:
		fmt.Println("放开上面的 fallthrough 才能看到我")
	}
}
