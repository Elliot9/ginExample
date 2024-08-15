package admin

import (
	//"errors"
	"fmt"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
	"strconv"
	// "github.com/go-playground/validator/v10"
)

type registerReqeust struct {
	Name     string `form:"name" binding:"required,max=20"`
	Account  string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,max=20,min=4"`
}

func (h *handler) Register() context.HandlerFunc {
	return func(c context.Context) {
		req := new(registerReqeust)

		if err := c.ShouldBindForm(req); err != nil {
			fmt.Println(err)
			c.Abort(http.StatusBadRequest, err)
			return
		}

		id, err := h.service.Register(req.Name, req.Account, req.Password)
		if err != nil {
			c.Abort(http.StatusBadRequest, err)
			return
		}

		c.JSON(200, map[string]string{
			"id": strconv.Itoa(id),
		})
	}
}
