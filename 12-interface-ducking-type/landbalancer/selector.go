package landbalancer

//Selector 挑选服务器节点
type Selector interface {
	Select() string
}

//DefaultBalancer 默认
type DefaultBalancer struct {
	//
}

// RandomBalancer 随机
type RandomBalancer struct {
	//
}

//Select 挑选
func (d *DefaultBalancer) Select() string {
	// 算法略
	return "host1:port1"
}

//Select 挑选
func (r *RandomBalancer) Select() string {
	// 算法略
	return "host2:port2"
}
