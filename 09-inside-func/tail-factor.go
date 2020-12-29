package main

func main() {
  tailFactor(9, 10)
}

func tailFactor(num, current uint64) uint64 {
	if num == 1 {
		return current
	}

	return tailFactor(num-1, num*current)
}
