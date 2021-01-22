package main

//go:noinline
func inc(x int) int {
  return x + 1
}

//go:noinline
func add(a, b int) int {
  return inc(a) + b
}

func main() {
  add(1, 2)
}
