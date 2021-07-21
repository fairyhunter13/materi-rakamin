package main

import "fmt"

type Notification interface {
	Send(destination, content string)
}

type Email struct {
	sender      string
	destination string
	subject     string
	body        string
}

func (e *Email) Send(destination, content string) {
	e.destination = destination
	e.body = content
	fmt.Printf("Send email to %s from %s, with subject: %s, body: %s.\n", e.destination, e.sender, e.subject, e.body)
}

type SMS struct {
	sender      string
	destination string
	text        string
}

func (s *SMS) Send(destination, content string) {
	s.destination = destination
	s.text = content
	fmt.Printf("Send sms to %s from %s, with text: %s.\n", s.destination, s.sender, s.text)
}

func abstraction() {

}
