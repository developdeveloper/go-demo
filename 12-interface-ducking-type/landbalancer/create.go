package landbalancer

//New 创建负载均衡器
func New(strategy string) Selector {
	switch strategy {
	case "random":
		return new(RandomBalancer)
	default:
		return new(DefaultBalancer)
	}
}
