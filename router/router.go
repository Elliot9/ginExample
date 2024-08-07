package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/elliot9/ginExample/repository/mysql"
	"github/elliot9/ginExample/repository/redis"
)

type Server struct {
	Mux   *gin.Engine
	Db    mysql.Repo
	Cache redis.Repo
}

func NewHTTPServer() (*Server, error) {
	mux := gin.New()
	mux.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"status": true,
			"from":   "Elliot",
		})
	})

	// 初始化 MYSQL
	dbRepo, err := mysql.New()
	if err != nil {
		// todo change to log
		fmt.Println(err)
	}

	// 初始化 Redis
	redisRepo, err := redis.New()
	if err != nil {
		// todo change to log
		fmt.Println(err)
	}

	return &Server{
		Mux:   mux,
		Db:    dbRepo,
		Cache: redisRepo,
	}, nil
}
