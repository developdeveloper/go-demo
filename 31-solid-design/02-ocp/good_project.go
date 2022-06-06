package main

import (
  "fmt"
)

type Workable interface {
  DoWork() bool 
}

type Programmer struct {}

func (p * Programmer) DoWork() bool {
  fmt.Println("coding...")
  return true
}

type Tester struct {}

func (t * Tester) DoWork() bool {
  fmt.Println("testing...")
  return false
}

type ProjectManagement struct {
  Workers []Workable
} 

//func NewProjectManagement() *ProjectManagement {
//  return &ProjectManagement{
//    Workers: []Workable {
//      &Programmer{},
//      &Tester{},
//    },
//  }
//}

func NewProjectManagement() *ProjectManagement {
  return &ProjectManagement{}
}

func (pm *ProjectManagement) AddWorker(workers ...Workable) {
  pm.Workers = append(pm.Workers, workers...)
}
  
func (pm *ProjectManagement) Process() bool {
  for _, worker := range pm.Workers {
    if !worker.DoWork() {
      return false
    }
  }
  
  return true 
}

// 增加发布者
type Publisher struct {}

func (p * Publisher) DoWork() bool {
  fmt.Println("publishing...")
  return true 
}

func main() {
  pm := NewProjectManagement()
  // pm.AddWorker(&Programmer{}, &Tester{})
  pm.AddWorker(&Programmer{}, &Tester{}, &Publisher{})
  
  if !pm.Process() {
    fmt.Println("Produce failed")
    return 
  }
  
  fmt.Println("Produce success")
}