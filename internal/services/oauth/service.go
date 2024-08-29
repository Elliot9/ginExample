package oauth

import (
	"github/elliot9/ginExample/internal/repository/mysql"
	"github/elliot9/ginExample/internal/repository/user"

	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
)

type Agent string

const (
	agentGoogle Agent = "google"
	agentFB     Agent = "facebook"
)

type Service interface {
	GetOauthConfig(agent Agent) *oauth2.Config
	GetQuery(agent Agent) string
	Callback(agent Agent, state, code string) (userInfo *UserInfo, err error)
	Login(userInfo *UserInfo) (accessToken, refreshToken string, err error)
}

type service struct {
	validator *validator.Validate
	userRepo  user.UserRepo
}

func New(db mysql.Repo, validator *validator.Validate) Service {
	return &service{
		validator: validator,
		userRepo:  user.New(db),
	}
}
