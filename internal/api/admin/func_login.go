package admin

import (
	"errors"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
)

type loginReqeust struct {
	Account  string `form:"username"`
	Password string `form:"password"`
}

func (h *handler) Login() context.HandlerFunc {
	return func(c context.Context) {
		req := new(loginReqeust)

		if err := c.ShouldBindForm(req); err != nil {
			c.Abort(http.StatusBadRequest, errors.New("bad Request"))
			return
		}

		c.JSON(200, map[string]string{
			"Account":  req.Account,
			"Password": req.Password,
		})
	}
}
