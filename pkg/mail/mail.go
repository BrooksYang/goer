package mail

import (
	"sync"
)

type Mailer struct {
	Driver Driver
}

var once sync.Once
var mailer *Mailer

func NewMailer() *Mailer {
	once.Do(func() {
		mailer = &Mailer{
			Driver: &SMTP{},
		}
	})

	return mailer
}

func (m *Mailer) Send(to string, subject string, content string) bool {
	return m.Driver.Send(to, subject, content)
}
