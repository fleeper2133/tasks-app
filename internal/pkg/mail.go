package pkg

import "net/smtp"

type MailManager interface {
	SendMessage(to []string, message []byte) error
}

type EmailConfig struct {
	From     string
	Password string
	SmtpHost string
	SmtpPort string
}

type SendMailManager struct {
	auth *smtp.Auth
	port string
	host string
	from string
}

func NewSendMailManager(emailConfig EmailConfig) *SendMailManager {
	auth := smtp.PlainAuth("", emailConfig.From, emailConfig.Password, emailConfig.SmtpHost)
	return &SendMailManager{
		auth: &auth,
		port: emailConfig.SmtpPort,
		host: emailConfig.SmtpHost,
		from: emailConfig.From,
	}
}

func (sm *SendMailManager) SendMessage(to []string, message []byte) error {
	if err := smtp.SendMail(sm.host+":"+sm.port, *sm.auth, sm.from, to, message); err != nil {
		return err
	}
	return nil
}
