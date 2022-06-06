package main

import (
  "fmt"
  "encoding/json"
)

type TextFormatter interface {
  TextFormat() string 
}

type JsonFormatter interface {
  JsonFormat() string 
}

type Person struct {
  Name string 
  Age uint8
}

func NewPerson(name string, age uint8) *Person {
  return &Person{Name: name, Age: age}
}

func (p *Person) Grow(years uint8) {
  p.Age += years
}

type PersonFormatter struct {
  P *Person
}

func (pf *PersonFormatter) TextFormat() string {
  return fmt.Sprintf("%s is %d old", pf.P.Name, pf.P.Age)
}

func (pf *PersonFormatter) JsonFormat() string {
  bytes, _ := json.Marshal(map[string]interface{} {
    "name": pf.P.Name,
    "age": pf.P.Age,
  })
  return string(bytes)
}

func printText(formatter TextFormatter) {
  str := formatter.TextFormat()
  fmt.Println(str)
}

func printJson(formatter JsonFormatter) {
  str := formatter.JsonFormat()
  fmt.Println(str)
}

func main() {
  person := NewPerson("zhangsan", 23)
  person.Grow(2)
  personFormatter := &PersonFormatter{P: person}
  printText(personFormatter) 
  printJson(personFormatter) 
}