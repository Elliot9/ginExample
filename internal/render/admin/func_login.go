package admin

import (
	"fmt"
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/context"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type loginReqeust struct {
	Account  string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,max=20,min=4"`
}

func (h *handler) Login() context.HandlerFunc {
	return func(c context.Context) {
		req := new(loginReqeust)

		if err := c.ShouldBindForm(req); err != nil {
			errors := make(map[string]any)
			for _, fieldErr := range err.(validator.ValidationErrors) {
				errors[fieldErr.Field()] = fieldErr.ActualTag()
			}

			c.ReturnBackWith(errors)
			return
		}

		admin, err := h.service.Login(req.Account, req.Password)
		if err != nil {
			c.ReturnBackWith(map[string]any{
				"Password": err.Error(),
			})
			return
		}

		if err := storeAuthSession(c, admin); err != nil {
			fmt.Println(err)
			c.Abort(context.Error(http.StatusInternalServerError, 500, "Store Session Error"))
			return
		}

		c.Redirect(http.StatusFound, "/admin")
	}
}

func storeAuthSession(c context.Context, admin *models.Admin) error {
	c.Session().Set(context.SessionAuthKey, &admin)

	if err := c.Session().Save(); err != nil {
		return err
	}

	return nil
}
