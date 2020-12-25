package main

import (
	"05-condition-loop/branch"
	"05-condition-loop/check"
	"05-condition-loop/loop"
)

func main() {
	check.CanBuy(80, 50)
	check.CheckPrice1(50)
	check.CheckPrice2(50)
	check.CheckPrice3(50)
	check.CheckPrice4(50)
	check.CheckPrice5(50)

	loop.FullFor()
	loop.ConditionFor()
	// loop.DeadFor()
	loop.BreakContinueFor()

	loop.ForRange()
	loop.ForRange2()
	loop.ForRange3()

	branch.FullSwitch("silver")
	branch.FullSwitch("nothing")
	branch.FallThrough()

	loop.JumpOutside()

	// branch.DeadSelect()
}
