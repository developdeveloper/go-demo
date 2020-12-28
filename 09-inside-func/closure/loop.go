package closure

import "fmt"

func loopForDelay() {
	for i := 0; i < 10; i++ {
		// defer func() {
		// 	fmt.Println(i) // 10 10 10 10 10 10 10 10 10 10
		// }()
		defer func(number int) {
			fmt.Println(number) // 9 8 7 6 5 4 3 2 1 0
		}(i)
	}
}
