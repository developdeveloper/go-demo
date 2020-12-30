package embed

type addr struct {
	address string
	phone   string
}

// type cat struct {
// 	color   string
// 	address string
// 	phone   string
// }

type cat struct {
	color string
	addr  // 匿名字段
}

// type person struct {
// 	sex     string
// 	address string
// 	phone   string
// }

type person struct {
	sex  string
	addr // 匿名字段
}
