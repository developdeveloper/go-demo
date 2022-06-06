package main

import (
  "fmt"
)

type Mailer interface {
  Body() string 
}

type TextMailer struct {}

func (m *TextMailer) Body() string {
  return "text message"
}

type HtmlMailer struct {}

func (m *HtmlMailer) Body() string {
  return "<span>html message</span>"
}

type WelcomeSender struct {
  // 别关心是什么 Mailer，我会给你注入进来，依赖被反转了
  M Mailer 
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