package main

import (
  "fmt"
)

func output(int) (int, int, int)

func main() {
  a, b, c := output(10001000)
  fmt.Println(a, b, c)
}
