package admin

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/pkg/context"

	"github.com/go-playground/validator/v10"
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
			errors := make(map[string]any)
			for _, fieldErr := range err.(validator.ValidationErrors) {
				errors[fieldErr.Field()] = fieldErr.ActualTag()
			}

			c.ReturnBackWith(errors)
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
			c.Abort(context.Error(http.StatusInternalServerError, 500, "Store Session Error"))
			return
		}

		c.Redirect(http.StatusFound, "/admin")
	}
}
