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

func (b *Buyer) BuyVet() bool { 
  fmt.Println(b.Drive() + " buying vet...")
  return true 
}

func (b *Buyer) BuyPig() bool {
  fmt.Println(b.Drive() + " buying pig...")
  return true   
}

func (b *Buyer) BuyEgg() bool { 
  fmt.Println(b.Drive() + " buying egg...")
  return true   
}

func (b *Buyer) BuyRice() bool { 
  fmt.Println(b.Drive() + " buying rice...")
  return true 
}

type Cooker struct {}

func (c *Cooker) Cooking(foods string) bool {
  if strings.Contains(foods, "红烧肉") {
    buyer := Buyer{&BikeCar{}}
    if !buyer.BuyVet() || !buyer.BuyPig() {
      return false 
    }
  } 
  
  if strings.Contains(foods, "蛋炒饭") {
    buyer := Buyer{&WalkCar{}}
    if !buyer.BuyEgg() || !buyer.BuyRice() {
      return false 
    }
  } 
  
  fmt.Println("cooking...")
  return true 
}

func main() {
  c := &Cooker {}
  c.Cooking("红烧肉 + 蛋炒饭")
}