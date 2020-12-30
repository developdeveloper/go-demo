package person

import "fmt"

//Person 人
type Person struct {
	Name      string
	Age       int
	WalkSpeed int
	SwimSpeed int
}

//Walk 走
func (p *Person) Walk(distance int) int {
	return 2 * distance / p.WalkSpeed
}

//Swim 游
func (p *Person) Swim(distance int) int {
	return 3 * distance / p.SwimSpeed
}

//Fly 飞行
func (p *Person) Fly() {
	fmt.Println("start fly now")
}
