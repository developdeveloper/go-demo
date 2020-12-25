package branch

// DeadSelect 一直阻塞
func DeadSelect() {
	select {}
}
