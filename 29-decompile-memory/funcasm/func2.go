package main

func add2(a, b int) int {
	times := 2
	println(times)
	return (a - b) * times
}

func main() {
	sum := add2(1, 2)
	println(sum)
}
