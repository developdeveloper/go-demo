package duck

import "fmt"

//Duck 鸭子
type Duck struct {
	Category  string
	Name      string
	WalkSpeed int
	SwimSpeed int
}

//Walk 走
func (d *Duck) Walk(distance int) int {
	return distance / d.WalkSpeed
}

//Swim 游
func (d *Duck) Swim(distance int) int {
	return distance / d.SwimSpeed
}

//Egg 下蛋
func (d *Duck) Egg() {
	fmt.Println("create egg now")
}
