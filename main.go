package main

import (
	"context"
	"encoding/gob"
	"errors"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/elliot9/gin-example/config"
	"github.com/elliot9/gin-example/internal/loader"
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/services/queue/consumer"
	"github.com/elliot9/gin-example/pkg/shutdown"

	"github.com/gin-gonic/gin"
)

func init() {
	// 註冊序列化類型
	gob.Register(map[string]interface{}{})
	gob.Register(models.Admin{})
	config.Load()
}

func main() {
	gin.SetMode(config.AppSetting.Env)
	s, err := loader.NewHTTPServer()

	if err != nil {
		panic(err)
	}

	defer func() {
		//_ = s.Db.DbWClose()
		//_ = s.Db.DbRClose()
		//_ = s.Cache.Close()
	}()

	server := &http.Server{
		Addr:    config.AppSetting.Url,
		Handler: s.Mux,
	}

	// 啟動 Http Server
	go func() {
		log.Printf("[info] Http Server start listening %s\n", config.AppSetting.Url)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server startup error: %v", err)
		}
		log.Println("[info] Stopped serving new connections.")
	}()

	// 啟動 AMQP Server
	amqpService := consumer.New(s.Amqp, s.Mailer)
	if amqpService == nil {
		log.Println("[info] AMQP server is nil")
	} else {
		go func() {
			if err := amqpService.EmailWelcome(); err != nil {
				log.Fatalf("AMQP server startup error: %v", err)
			}

			log.Println("[info] AMQP Server stop listening.")
		}()
	}

	// 順序關閉 Http Server、AMQP Server、Redis、DB
	shutdown.New(syscall.SIGINT, syscall.SIGTERM).OnShutdown(func() {
		// 關閉 Http Server
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("HTTP shutdown error: %v", err)
		}
		log.Println("[info] Http Server shutdown complete.")
	}, func() {
		// 關閉 AMQP connection
		if s.Amqp != nil {
			_ = s.Amqp.Close()
			log.Println("[info] AMQP shutdown complete.")
		}
	}, func() {
		// 關閉 Redis connection
		if s.Cache != nil {
			_ = s.Cache.Close()
			log.Println("[info] Redis shutdown complete.")
		}
	}, func() {
		if s.Db == nil {
			return
		}
		// 關閉 DB connection
		if s.Db.GetDbR() != nil {
			_ = s.Db.DbRClose()
		}

		if s.Db.GetDbW() != nil {
			_ = s.Db.DbWClose()
		}
		log.Println("[info] DB shutdown complete.")
	})
}
