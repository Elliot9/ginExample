package admin

import (
	//"errors"
	"fmt"
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"
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
			c.Abort(context.Error(http.StatusBadRequest, 100, err.Error()))
			return
		}

		id, err := h.service.Register(req.Name, req.Account, req.Password)
		if err != nil {
			c.Abort(context.Error(http.StatusBadRequest, 100, err.Error()))
			return
		}

		admin := &models.Admin{Model: models.Model{ID: uint(id)}, Name: req.Name, Email: req.Account}

		if err := storeAuthSession(c, admin); err != nil {
			fmt.Println(err)
			c.Abort(context.Error(http.StatusInternalServerError, 500, "儲存 Session Error"))
			return
		}

		c.Redirect(http.StatusFound, "/admin")
	}
}
