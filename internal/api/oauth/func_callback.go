package oauth

import (
	"fmt"

	"net/http"

	"github.com/elliot9/gin-example/internal/pkg/context"
	"github.com/elliot9/gin-example/internal/services/oauth"
)

type getCallbackRequest struct {
	State string `form:"state"`
	Code  string `form:"code"`
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
		if err := c.ShouldBindForm(req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		userInfo, err := h.service.Callback(oauth.Agent(agent.Agent), req.State, req.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Printf("userInfo: %v\n", userInfo)
		accessToken, refreshToken, err := h.service.Login(userInfo)
		fmt.Printf("accessToken: %v\n", accessToken)
		fmt.Printf("refreshToken: %v\n", refreshToken)
		fmt.Printf("err: %v\n", err)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	}
}
