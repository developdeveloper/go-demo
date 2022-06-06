package main

import (
  "fmt"
)

type TextMailer struct {}

func (m *TextMailer) Body() string {
  return "text message"
}

type HtmlMailer struct {}

func (m *HtmlMailer) Body() string {
  return "<span>html message</span>"
}

type WelcomeSender struct {
  // 需要关系是哪种类型的 Mailer，需要直接构建，依赖外部具体的类型
  M *TextMailer
  // *M HtmlMailer
}

func (ws *WelcomeSender) DoWelcome() {
  str := ws.M.Body()
  fmt.Println(str)
}

func main() {
  ws := &WelcomeSender{&TextMailer{}}
  ws.DoWelcome()
  
  ws = &WelcomeSender{&HtmlMailer{}}
  ws.DoWelcome()
}