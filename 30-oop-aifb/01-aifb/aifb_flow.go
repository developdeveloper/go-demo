package main

// a = abstract     抽象
// i = implmention  实现
// f = framework    框架
// b = business     业务

import (
  "fmt"
)

// 抽象
type Phone interface {
  Call(name string) 
}

// 实现

// Huawei
type HuaweiPhone struct {} 

func (phone *HuaweiPhone) Call(name string) {
  fmt.Println("huawei call " + name)
}

// Apple 
type ApplePhone struct {} 

func (phone *ApplePhone) Call(name string) {
  fmt.Println("apple call " + name)
}

// 框架
func doCall(phone Phone, names... string) {
  for _, name := range names {
    phone.Call(name)
  }
}

// 扩展
// OppoPhone 
type OppoPhone struct { }

func (phone *OppoPhone) Call(name string) {
  fmt.Println("oppo call " + name)
}

// 业务
func main() {
  var phone Phone
  
  phone = &ApplePhone{}
  doCall(phone, "zhangsan", "lisi")
  
  phone = &HuaweiPhone{}
  doCall(phone, "zhangsan", "lisi")
}