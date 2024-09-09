package mail

import (
	"bytes"
	"html/template"
	"os"

	"github.com/elliot9/gin-example/config"
)

func (s *service) Welcome(to, userName, verificationLink string) error {
	//讀取文件內容
	content, err := os.ReadFile("internal/assets/templates/mail/welcome.html")
	if err != nil {
		return err
	}

	// 解析模板
	tmpl, err := template.New("email").Parse(string(content))
	if err != nil {
		return err
	}

	// 渲染模板
	var body bytes.Buffer
	data := map[string]string{
		"Username":         userName,
		"VerificationLink": verificationLink,
		"TeamName":         config.AppSetting.Name,
		"CompanyName":      config.AppSetting.Name,
	}

	if err := tmpl.ExecuteTemplate(&body, "mail/welcome", data); err != nil {
		return err
	}

	return s.mailer.SendHTML([]string{to}, "Welcome to "+config.AppSetting.Name, &body)
}
