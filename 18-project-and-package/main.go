package main

import (
	passwordGenerator "github.com/developdeveloper/password-generator/pkg"
	"github.com/fatih/color"
)

func main() {
	//color.Cyan("我是 cyan")
	//color.Red("我是 red")
	pwdGen := passwordGenerator.New()
	color.Green(pwdGen.DigitsOnly(6))
	color.Blue(pwdGen.CharsOnly(6))
}
