package main

import (
	"context"
	"encoding/gob"
	"errors"
	"github/elliot9/ginExample/config"
	"github/elliot9/ginExample/internal/loader"
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/pkg/shutdown"
	"log"
	"net/http"
	"syscall"
	"time"

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

	go func() {
		log.Printf("[info] Http Server start listening %s\n", config.AppSetting.Url)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server startup error: %v", err)
		}
		log.Println("[info] Stopped serving new connections.")
	}()

	shutdown.New(syscall.SIGINT, syscall.SIGTERM).OnShutdown(func() {
		// 關閉 Http Server
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("HTTP shutdown error: %v", err)
		}
		log.Println("[info] Http Server shutdown complete.")
	}, func() {
		// 關閉 DB connection
		if s.Db.GetDbR() != nil {
			_ = s.Db.DbRClose()
		}

		if s.Db.GetDbW() != nil {
			_ = s.Db.DbWClose()
		}
		log.Println("[info] DB shutdown complete.")
	}, func() {
		// 關閉 Redis connection
		if s.Cache != nil {
			_ = s.Cache.Close()
			log.Println("[info] Redis shutdown complete.")
		}
	})
}
