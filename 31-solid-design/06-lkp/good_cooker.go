package main

import (
  "fmt"
  "strings"
)

type Runable interface {
  Drive() string
}

type BikeCar struct {}

func (bc * BikeCar) Drive() string {
  return "car..."
}

type WalkCar struct {}

func (wc * WalkCar) Drive() string {
  return "walk..."
}

type Buyer struct {
  Runable
}

func (b *Buyer) buyVet() bool { 
  fmt.Println(b.Drive() + " buying vet...")
  return true 
}

func (b *Buyer) buyPig() bool {
  fmt.Println(b.Drive() + " buying pig...")
  return true   
}

func (b *Buyer) buyEgg() bool { 
  fmt.Println(b.Drive() + " buying egg...")
  return true 
}

func (b *Buyer) buyRice() bool { 
  fmt.Println(b.Drive() + " buying rice...")
  return true 
}

func (b* Buyer) DoFirePig() bool {
  if b.buyVet() && b.buyPig() {
    return true 
  } 
  
  return false 
}

func (b* Buyer) DoFireEgg() bool {
  if b.buyEgg() && b.buyRice() {
    return true 
  } 
  
  return false 
}

func (b *Buyer) BuyMaterial(foods string) bool {
  if strings.Contains(foods, "红烧肉") {
    buyer := Buyer{&BikeCar{}}
    if !buyer.DoFirePig() {
      return false 
    }
  }
  
  if strings.Contains(foods, "蛋炒饭") {
    buyer := Buyer{&WalkCar{}}
    if !buyer.DoFireEgg() {
      return false  
    }
  }
  
  return true 
}

type Cooker struct {}

func (c *Cooker) Cooking(foods string) {
  buyer := Buyer {}
  
  if buyer.BuyMaterial(foods) {
    fmt.Println("cooking...")
  }
}

func main() {
  c := &Cooker {}
  c.Cooking("红烧肉 + 蛋炒饭")
}