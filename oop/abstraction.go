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

type NotificationSender struct {
	notifications map[string]Notification
}

type NotifParams struct {
	Destination string
	Content     string
}

func (np *NotifParams) IsValid() bool {
	return np.Content != "" && np.Destination != ""
}

func (ns *NotificationSender) Send(medium string, params NotifParams) {
	if !params.IsValid() {
		return
	}

	notif, ok := ns.notifications[medium]
	if !ok {
		return
	}

	notif.Send(params.Destination, params.Content)
}

func abstraction() {
	ns := NotificationSender{
		notifications: map[string]Notification{
			"email": &Email{
				sender:  "me@mail.com",
				subject: "no-reply",
			},
			"sms": &SMS{
				sender: "0812109213123",
			},
		},
	}

	ns.Send("email", NotifParams{
		Destination: "our@mail.com",
		Content:     "This is an email",
	})

	ns.Send("sms", NotifParams{
		Destination: "0811112223334",
		Content:     "This is an sms",
	})

	ns.Send("email", NotifParams{
		Content: "This doesn't get sent",
	})
}
