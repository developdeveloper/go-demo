package enum

import "fmt"

// const (
// 	appple = 0
// 	banana = 1
// 	orange = 2
// )

// const (
// 	appple = iota
// 	banana = iota
// 	orange = iota
// )

const (
	appple = iota
	banana
	orange
)

//ApplyEnum 当做枚举
func ApplyEnum() {
	fmt.Println(appple, banana, orange)
}
