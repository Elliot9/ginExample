package config

type Mailer struct {
	Host       string
	Port       int
	UserName   string
	Password   string
	SenderName string
}

var MailerSetting *Mailer
