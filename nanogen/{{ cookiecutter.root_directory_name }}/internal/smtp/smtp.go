package smtp

import (
	"function/internal/config"
	"log"

	"gopkg.in/gomail.v2"
)

type SMTP struct {
	Host     *string
	Port     *int
	Username *string
	Password *string
	Sender   *string
}

func New(cfg *config.SmtpConfig) *SMTP {
	return &SMTP{
		Host:     &cfg.Host,
		Port:     &cfg.Port,
		Username: &cfg.Username,
		Password: &cfg.Password,
		Sender:   &cfg.Sender,
	}
}

func (s *SMTP) SendEmail(receiver string, subject string, body string) (string, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", *s.Sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := &gomail.Dialer{Host: *s.Host, Port: *s.Port}

	if len(*s.Username) > 0 && len(*s.Password) > 0 {
		d.Username = *s.Username
		d.Password = *s.Password
	}

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email [%s]:\n", err)
		return "", err
	}

	return subject + " email sent", nil
}
