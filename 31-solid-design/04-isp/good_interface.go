package main

import (
  "fmt"
)

type Codable interface {
  Code() bool 
}

type Testable interface {
  Test() bool 
}

type Programmer struct {}

func (p *Programmer) Code() bool {
  fmt.Println("coding...")
  return true 
}

func (p *Programmer) Test() bool {
  fmt.Println("testing on local...")
  return true 
}

type Tester struct {}

func (t *Tester) Test() bool {
  fmt.Println("Testing on server...")
  return true
}

type ProjectManagement struct {
  Workers []interface{}
}

func (pm *ProjectManagement) AddWorker(worker interface{}) {
  pm.Workers = append(pm.Workers, worker)
}

func (pm *ProjectManagement) Produce() {
  for _, worker := range pm.Workers {
    if w, ok := worker.(Codable); ok {
      w.Code()
    }
    
    if w, ok := worker.(Testable); ok {
      w.Test()
    }
  }
}

func main() {
  pm := ProjectManagement{}
  pm.AddWorker(&Programmer{})
  pm.AddWorker(&Tester{})
  pm.Produce()
}