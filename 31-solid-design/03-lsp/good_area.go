package main

import (
  "fmt"
)

type Areable interface {
  Area() float32
}

type Rectangle struct {
  Width float32
  Height float32
}

func NewRectangle(width, height float32) *Rectangle {
  return &Rectangle{Width: width, Height: height}
}

func (r *Rectangle) Area() float32 {
  return r.Width * r.Height
}

type Square struct {
  Rectangle
}

func NewSquare(side float32) *Square {
  return &Square{Rectangle{Width: side, Height: side}}
}

func printArea(area Areable) {
  fmt.Println(area.Area())
}

func main() {
  r := &Rectangle{Width: 2, Height: 4}
  printArea(r)

  s := NewSquare(2)
  printArea(s)
}
