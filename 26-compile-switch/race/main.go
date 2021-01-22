package main

var sum int

func main() {
  go add()
  go add()
}

func add() {
  sum++
}
