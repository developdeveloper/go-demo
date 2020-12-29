package main

func main() {
  factor2(10)
}

func factor2(num uint64) uint64 {
	if num == 1 {
		return 1
	}

	return num * factor2(num-1)
}
