package health

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"time"
)

type Handler interface {
	Ping() context.HandlerFunc
}

type handler struct {
}

func (h *handler) Ping() context.HandlerFunc {
	return func(c context.Context) {
		resp := &struct {
			Timestamp   time.Time `json:"timestamp"`
			Environment string    `json:"environment"`
			Host        string    `json:"host"`
			Status      string    `json:"status"`
		}{
			Timestamp:   time.Now(),
			Environment: "Testing",
			Host:        "host",
			Status:      "ok",
		}

		c.JSON(200, resp)
	}
}

func New() *handler {
	return &handler{}
}
