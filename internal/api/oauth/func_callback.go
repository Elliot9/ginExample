package oauth

import (
	"github/elliot9/ginExample/internal/pkg/context"
	"github/elliot9/ginExample/internal/services/oauth"
	"net/http"
)

type getCallbackRequest struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

func (h *handler) Callback() context.HandlerFunc {
	return func(c context.Context) {
		var agent struct {
			Agent string `uri:"agent"`
		}
		if err := c.ShouldBindURI(&agent); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		req := new(getCallbackRequest)
		if err := c.ShouldBindJson(req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := h.service.Callback(oauth.Agent(agent.Agent), req.State, req.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"user": user.Email,
		})
	}
}
