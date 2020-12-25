package main

import (
	"06-const-var-reference/constant"
	"06-const-var-reference/enum"
	"06-const-var-reference/varpointer"
)

func main() {
	constant.CompileSizedof()

	enum.ApplyEnum()

	varpointer.PrintAddr()
	varpointer.HoldVarAddr()
	varpointer.IncScoreFailed()
	varpointer.IncScoreSuccssed()
}
