package loader

import (
	"fmt"
	"github/elliot9/ginExample/config"
	"github/elliot9/ginExample/internal/repository/amqp"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"
	"github/elliot9/ginExample/internal/router"
	"github/elliot9/ginExample/pkg/mailer"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	Mux       *gin.Engine
	Validator *validator.Validate
	Db        mysql.Repo
	Cache     redis.Repo
	Mailer    mailer.Mailer
	Amqp      amqp.Repo
}

func NewHTTPServer() (*Server, error) {
	mux := gin.New()
	// 初始化 MYSQL
	dbRepo, err := mysql.New()
	if err != nil {
		// todo change to log
		fmt.Println(err)
	}
	log.Println("[info] DB connection")

	// 初始化 Redis
	redisRepo, err := redis.New()
	if err != nil {
		// todo change to log
		fmt.Println(err)
	}
	log.Println("[info] Redis connection")

	// 初始化 Validator
	validator := validator.New()

	// 初始化 Mailer
	mailer := mailer.New(getMailConfig())

	// 初始化 Amqp
	amqp, err := amqp.New()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("[info] Amqp connection")

	// 註冊 Router
	router.RegisterRouter(mux, dbRepo, redisRepo, validator, mailer, amqp)

	return &Server{
		Mux:       mux,
		Validator: validator,
		Db:        dbRepo,
		Cache:     redisRepo,
		Mailer:    mailer,
		Amqp:      amqp,
	}, nil
}

func getMailConfig() *mailer.Option {
	return &mailer.Option{
		Host:       config.MailerSetting.Host,
		Port:       config.MailerSetting.Port,
		User:       config.MailerSetting.UserName,
		Password:   config.MailerSetting.Password,
		SenderName: config.MailerSetting.SenderName,
	}
}
