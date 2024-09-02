package amqp

import (
	"context"
	"fmt"
	"github/elliot9/ginExample/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Repo interface {
	Close() error
	QueueDeclare(name string, options *QueueDeclareOptions) (amqp.Queue, error)
	Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Consume(ctx context.Context, queue *amqp.Queue, fn func(msg amqp.Delivery) error, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) error
}

type QueueDeclareOptions struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type repo struct {
	conn *amqp.Connection
}

func New() (Repo, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.AmqpSetting.User, config.AmqpSetting.Password, config.AmqpSetting.Host, config.AmqpSetting.Port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return &repo{conn: conn}, nil
}

func (r *repo) Close() error {
	return r.conn.Close()
}

func (r *repo) QueueDeclare(name string, options *QueueDeclareOptions) (amqp.Queue, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	if options == nil {
		options = &QueueDeclareOptions{}
	}

	return ch.QueueDeclare(name, options.Durable, options.AutoDelete, options.Exclusive, options.NoWait, options.Args)
}

func (r *repo) Publish(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	return ch.PublishWithContext(ctx,
		exchange,
		key,
		mandatory,
		immediate,
		msg,
	)
}

func (r *repo) Consume(ctx context.Context, queue *amqp.Queue, fn func(msg amqp.Delivery) error, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queue.Name,
		"",
		autoAck,
		exclusive,
		noLocal,
		noWait,
		args,
	)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %w", err)
	}

	var forever chan struct{}
	go func() {
		for msg := range msgs {
			var deliveryCount int64
			if msg.Headers["x-delivery-count"] != nil {
				deliveryCount = msg.Headers["x-delivery-count"].(int64)
			}
			if deliveryCount > 3 {
				log.Printf("Delivery count exceeded: %d\n", deliveryCount)
				msg.Nack(false, false)
				continue
			}

			log.Printf("Received a message: %s\n", msg.Body)
			if err := fn(msg); err != nil {
				log.Printf("Error processing message: %v", err)
				msg.Nack(false, true)
			} else {
				msg.Ack(false)
			}
		}
	}()
	<-forever
	return nil
}
