package auth

import (
	"github.com/elliot9/gin-example/internal/models"
	"github.com/elliot9/gin-example/internal/pkg/context"
)

type Auth interface {
	Me(ctx context.Context) *models.Admin
}

var _ Auth = (*auth)(nil)

func New() Auth { return &auth{} }

type auth struct{}

func (a *auth) Me(ctx context.Context) *models.Admin {
	admin := ctx.Session().Get(context.SessionAuthKey)
	if admin == nil {
		return nil
	}

	adminModel, ok := admin.(models.Admin)
	if !ok {
		return nil
	}
	return &adminModel
}
