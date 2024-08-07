package main

import (
	"github/elliot9/ginExample/config"
	"github/elliot9/ginExample/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
}

func main() {
	gin.SetMode(config.AppSetting.Env)
	s, err := router.NewHTTPServer()

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = s.Db.DbWClose()
		_ = s.Db.DbRClose()
		_ = s.Cache.Close()
	}()

	server := &http.Server{
		Addr:    config.AppSetting.Url,
		Handler: s.Mux,
	}

	log.Printf("[info] start http server listening %s", config.AppSetting.Url)
	_ = server.ListenAndServe()
}
