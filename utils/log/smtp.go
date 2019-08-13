package log

import (
	"fmt"
	"net/smtp"
	"strings"
)

type SMTPConfig struct {
	Enable          bool   `json:"enable" env:"SMTP_ENABLE"`
	Addr            string `json:"addr" env:"SMTP_ADDR"`
	Host            string
	Username        string   `json:"username" env:"SMTP_USERNAME"`
	Password        string   `json:"password" env:"SMTP_PASSWORD"`
	MailTo          []string `json:"to" env:"SMTP_MAILTO"`
	MailFrom        string   `json:"from" env:"SMTP_MAILFROM" envDefault:"sausageshooter@xindong.com"`
	ContentType     string   `json:"content_type" env:"SMTP_CONTENT_TYPE" envDefault:"Content-Type: text/plain; charset=UTF-8"`
	MessageTemplate string   `json:"message_template" env:"SMTP_MSG_TEMPLATE" envDefault:"To:%s\r\nFrom:<%s>\r\nSubject:Sausage Panic Log\r\n%s\r\n\r\n%s"`

	mailToString string
}

var config = new(SMTPConfig)

func InitSMTP(smtpConfig SMTPConfig) {
	config = &smtpConfig
	config.Host = strings.Split(config.Addr, ":")[0]
	config.mailToString = strings.Join(config.MailTo, ";")
}

func SendMail(body string) error {
	if !config.Enable {
		return nil
	}
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	msg := fmt.Sprintf(config.MessageTemplate, config.mailToString, config.MailFrom, config.ContentType, body)

	return smtp.SendMail(config.Addr, auth, config.MailFrom, config.MailTo, []byte(msg))
}
