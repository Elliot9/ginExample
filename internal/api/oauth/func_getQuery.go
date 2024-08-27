package oauth

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/services/oauth"
	"net/http"
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
