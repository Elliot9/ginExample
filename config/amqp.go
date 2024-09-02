package config

type Amqp struct {
	Host     string
	Port     int
	User     string
	Password string
}

var AmqpSetting *Amqp
