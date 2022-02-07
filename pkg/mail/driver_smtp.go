package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"goer/global"

	"github.com/jordan-wright/email"
)

type SMTP struct{}

func (s *SMTP) Send(to string, subject string, content string) bool {
	// Get email config
	emailCfg := global.Config.Mail

	// New instance
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", emailCfg.FromName, emailCfg.FromAddress)
	e.To = []string{to}
	e.Bcc = []string{}
	e.Cc = []string{}
	e.Subject = subject
	e.Text = []byte(content)

	// Addr & Auth
	addr := fmt.Sprintf("%s:%d", emailCfg.Host, emailCfg.Port)
	auth := smtp.PlainAuth("", emailCfg.Username, emailCfg.Password, emailCfg.Host)

	// Send Email
	err := e.SendWithTLS(addr, auth, &tls.Config{ServerName: emailCfg.Host})
	if err != nil {
		global.Logger.Mail.Error("email failed, err: " + err.Error())
		return false
	}

	global.Logger.Mail.Info(fmt.Sprintf("email sent, to: %s, subject: %s, content: %s", to, subject, content))

	return true
}
