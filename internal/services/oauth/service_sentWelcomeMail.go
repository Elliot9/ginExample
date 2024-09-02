package oauth

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *service) SentWelcomeMail(to, name, url string) error {
	data := map[string]string{
		"to":   to,
		"name": name,
		"url":  url,
	}
	jsonData, _ := json.Marshal(data)

	err := s.amqp.Publish(context.Background(), "", "email_welcome", false, false, amqp.Publishing{
		Body: jsonData,
	})

	if err != nil {
		return fmt.Errorf("發歡迎郵件失敗: %w", err)
	}

	return nil
}
