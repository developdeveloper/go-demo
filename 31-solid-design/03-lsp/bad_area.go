package main

import (
  "fmt"
)

type Rectangle struct {
  Width float32
  Height float32
}

func (r *Rectangle) Area() float32 {
  return r.Width * r.Height
}

func NewRectangle(width, height float32) *Rectangle {
  return &Rectangle{Width: width, Height: height}
}

type Square struct {
  Rectangle
}

func NewSquare(side float32) *Square {
  return &Square{Rectangle{Width: side, Height: side}}
}

func printArea(rect *Rectangle) {
  fmt.Println(rect.Area())
}

func main() {
  r := NewRectangle(2, 3)
  printArea(r)

  s := NewSquare(2)
  printArea(s)
}
