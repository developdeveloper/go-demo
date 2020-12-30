package main

import (
	"12-interface-ducking-type/duck"
	"12-interface-ducking-type/move"
	"12-interface-ducking-type/person"
)

func hurryUp(moveable move.Movable, distance int) int {
	return moveable.Walk(distance/2) +
		moveable.Swim(distance/2)
}

func main() {
	d := &duck.Duck{Name: "kity", WalkSpeed: 10, SwimSpeed: 20}
	p := &person.Person{Name: "zhangsan", WalkSpeed: 20, SwimSpeed: 10}

	hurryUp(d, 2000)
	hurryUp(p, 2000)

	d.Egg()
	p.Fly()
}
