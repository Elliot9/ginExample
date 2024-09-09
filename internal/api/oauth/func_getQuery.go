package oauth

import (
	"net/http"

	"github.com/elliot9/gin-example/internal/pkg/context"
	"github.com/elliot9/gin-example/internal/services/oauth"
)

type getQueryRequest struct {
	Agent string `uri:"agent"`
}

func (h *handler) GetQuery() context.HandlerFunc {
	return func(c context.Context) {
		req := new(getQueryRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		agent := req.Agent
		url := h.service.GetQuery(oauth.Agent(agent))

		c.JSON(http.StatusOK, map[string]string{
			"url": url,
		})
	}
}
