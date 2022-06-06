package main

import (
  "fmt"
)

type Workable interface {
  Code() bool 
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

func (t *Tester) Code() bool {
  // do nothing
  return true
}

func (t *Tester) Test() bool {
  fmt.Println("Testing on server...")
  return true
}

type ProjectManagement struct {
  Workers []Workable
}

func (pm *ProjectManagement) AddWorker(worker Workable) {
  pm.Workers = append(pm.Workers, worker)
}

func (pm *ProjectManagement) Produce() {
  for _, worker := range pm.Workers {
    worker.Code()
    worker.Test()
  }
}

func main() {
  pm := ProjectManagement{}
  pm.AddWorker(&Programmer{})
  pm.AddWorker(&Tester{})
  pm.Produce()
}