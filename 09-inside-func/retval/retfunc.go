package retval

func getNameFunc(name string) func() string {
	return func() string {
		return name
	}
}
