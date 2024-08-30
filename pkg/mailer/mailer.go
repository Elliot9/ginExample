package mailer

import (
	"bytes"
	"net/smtp"
	"strconv"
	"strings"
)

type Mailer interface {
	Send(to []string, subject string, body string) error
	SendHTML(to []string, subject string, body *bytes.Buffer) error
	Option(option Option) Mailer
}

type Option struct {
	Host       string
	Port       int
	User       string
	Password   string
	SenderName string
}

var defaultOption = Option{
	Host:       "smtp.gmail.com",
	Port:       587,
	User:       "your-email@gmail.com",
	Password:   "your-password",
	SenderName: "Mailer",
}

type mailerImpl struct {
	option Option
}

func New(option *Option) Mailer {
	m := &mailerImpl{}
	if option == nil {
		m.Option(defaultOption)
	} else {
		m.Option(*option)
	}
	return m
}

func (m *mailerImpl) Option(option Option) Mailer {
	m.option = option
	return m
}

func (m *mailerImpl) Send(to []string, subject string, body string) error {
	msg := []byte("From: " + m.option.SenderName + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	return smtp.SendMail(m.option.Host+":"+strconv.Itoa(m.option.Port),
		smtp.PlainAuth("", m.option.User, m.option.Password, m.option.Host),
		m.option.User, to, msg)
}

func (m *mailerImpl) SendHTML(to []string, subject string, body *bytes.Buffer) error {
	msg := []byte("From: " + m.option.SenderName + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body.String())

	return smtp.SendMail(m.option.Host+":"+strconv.Itoa(m.option.Port),
		smtp.PlainAuth("", m.option.User, m.option.Password, m.option.Host),
		m.option.User, to, msg)
}
