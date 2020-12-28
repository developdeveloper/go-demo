package recfunc

func factor(num uint64) (res uint64) {
	if num > 1 {
		res = num * factor(num-1)
		return
	}

	return 1
}

func factor2(num uint64) uint64 {
	if num == 1 {
		return 1
	}

	return num * factor2(num-1)
}

func tailFactor(num, current uint64) uint64 {
	if num == 1 {
		return current
	}

	return tailFactor(num-1, num*current)
}
