package main

import (
  "fmt"
)

// 程序员
type Programmer struct {}

func (p * Programmer) DoCoding() bool {
  fmt.Println("coding...")
  return false 
}

// 测试员
type Tester struct {}

func (t * Tester) DoTesting() bool {
  fmt.Println("testing...")
  return false 
}

//// ①增加发布者修改
//type Publisher struct {}
//
//func (p * Publisher) DoWork() bool {
//  fmt.Println("publishing...")
//  return true 
//}

// 工程管理
type ProjectManagement struct {
  P *Programmer
  T *Tester
}

func NewProjectManagement() *ProjectManagement {
  // ②增加发布者修改
  return &ProjectManagement{
    P: &Programmer{},
    T: &Tester{},
  }
}

func (pm *ProjectManagement) Process() bool {
  // ③增加发布者修改 
  if pm.P.DoCoding() && pm.T.DoTesting() {
    return true 
  }
  
  return false 
}

func main() {
  pm := NewProjectManagement()
  
  if !pm.Process() {
    fmt.Println("Produce failed")
    return 
  }
  
  fmt.Println("Produce success")
}