package loop

import "fmt"

//JumpOutside 使用 goto 跳出多层循环
func JumpOutside() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				fmt.Println(i, j, k)
				if i == 1 && j == 2 && k == 3 {
					goto GAMEOVER
				}
			}
		}
	}

GAMEOVER:
	fmt.Println("game over")
}
