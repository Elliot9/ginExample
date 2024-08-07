package config

type Redis struct {
	Host     string
	Port     int
	Password string
}

var RedisSetting *Redis
