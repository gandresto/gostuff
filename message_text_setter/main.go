package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func (m *Message) SetText(text string) {
	m.Text = text
}

func main() {
	m := Message{
		Recipient: "Andres",
		Text:      "Hello World!",
	}
	fmt.Println(m.Text)
	m.SetText("Another text")
	fmt.Println(m.Text)
}
