package loader

import (
	"fmt"
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/redis"
	"github/elliot9/ginExample/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Mux   *gin.Engine
	Db    mysql.Repo
	Cache redis.Repo
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

	// 註冊 Router
	router.RegisterRouter(mux, dbRepo, redisRepo)

	return &Server{
		Mux:   mux,
		Db:    dbRepo,
		Cache: redisRepo,
	}, nil
}
