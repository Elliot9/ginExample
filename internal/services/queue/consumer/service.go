package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github/elliot9/ginExample/internal/repository/amqp"
	"github/elliot9/ginExample/internal/services/mail"
	"github/elliot9/ginExample/pkg/mailer"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	EmailWelcome() error
}

type service struct {
	amqp   amqp.Repo
	mailer mailer.Mailer
}

func New(amqp amqp.Repo, mailer mailer.Mailer) Service {
	return &service{
		amqp:   amqp,
		mailer: mailer,
	}
}

func (s *service) EmailWelcome() error {
	queue, err := s.amqp.QueueDeclare("email_welcome", &amqp.QueueDeclareOptions{
		Durable: true,
		Args:    amqp091.Table{"x-queue-type": "quorum"},
	})
	if err != nil {
		log.Fatalf("AMQP server startup error: %v", err)
	}

	return s.amqp.Consume(
		context.Background(),
		&queue,
		func(msg amqp091.Delivery) error {
			body := msg.Body
			var data map[string]string
			err := json.Unmarshal(body, &data)
			if err != nil {
				return fmt.Errorf("unmarshal error: %w", err)
			}

			to := data["to"]
			name := data["name"]
			url := data["url"]

			return mail.New(s.mailer).Welcome(to, name, url)
		},
		false,
		false,
		false,
		false,
		nil,
	)
}
