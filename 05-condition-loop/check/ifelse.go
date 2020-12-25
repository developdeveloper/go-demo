package check

import "fmt"

//CanBuy 判断钱是否足够
func CanBuy(hasMoney, productPrice int) {
	// if 后面不需要括号，if 的执行体无论是否为空 { } 都是必须的
	if hasMoney >= productPrice {
		fmt.Printf("可以买，还剩 %d 元\n", hasMoney-productPrice)
	} else {
		fmt.Println("开玩笑吗?")
	}
}

//CheckPrice1 判断贵不贵
func CheckPrice1(amount int) {
	if amount > 50 {
		if amount > 100 {
			fmt.Println("太贵了")
		} else {
			fmt.Println("还行")
		}
	} else {
		fmt.Println("真便宜啊")
	}
}

//CheckPrice2 判断贵不贵
func CheckPrice2(amount int) {
	if amount > 50 {
		if amount > 100 {
			fmt.Println("太贵了")
		} else {
			fmt.Println("还行")
		}
		return
	}

	fmt.Println("真便宜啊")
}

//CheckPrice3 判断贵不贵
func CheckPrice3(amount int) {
	if amount > 100 {
		fmt.Println("太贵了")
	} else if amount > 50 {
		fmt.Println("还行")
	} else {
		fmt.Println("真便宜啊")
	}
}

//CheckPrice4 判断贵不贵
func CheckPrice4(amount int) {
	if amount <= 50 {
		fmt.Println("真便宜啊")
		return
	}

	if amount > 50 && amount <= 100 {
		fmt.Println("还行")
	} else {
		fmt.Println("太贵了")
	}
}

//CheckPrice5 判断贵不贵
func CheckPrice5(amount int) {
	if amount > 100 {
		fmt.Println("太贵了")
		return
	}

	if amount > 50 {
		fmt.Println("还行")
		return
	}

	fmt.Println("真便宜啊")
}
