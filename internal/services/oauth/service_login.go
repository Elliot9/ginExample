package oauth

import (
	"github/elliot9/ginExample/internal/models"
	"github/elliot9/ginExample/internal/pkg/jwt"
	"time"

	"gorm.io/gorm"
)

func (s *service) Login(userInfo *UserInfo) (accessToken, refreshToken string, err error) {
	user, err := s.userRepo.FindByEmail(userInfo.Email)

	if err != nil && err != gorm.ErrRecordNotFound {
		return "", "", err
	}

	if user == nil {
		user = &models.User{
			Email:   userInfo.Email,
			Name:    userInfo.Name,
			Agent:   string(userInfo.Agent),
			AgentID: userInfo.Sub,
			State:   false,
		}
		user, err = s.userRepo.Create(user)
		if err != nil {
			return "", "", err
		}

		err = s.SentWelcomeMail(userInfo.Email, userInfo.Name, "https://example.com/verify")
		if err != nil {
			return "", "", err
		}
	}

	refreshToken = jwt.GenerateRefreshToken()
	accessToken, err = jwt.GenerateToken(user.Email, map[string]any{
		"name": user.Name,
	})

	if err != nil {
		return "", "", err
	}

	userRefreshToken := &models.UserRefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiredAt: time.Now().Add(time.Hour * 24 * 7),
	}

	err = s.userRepo.UpdateRefreshToken(userRefreshToken)
	if err != nil {
		return "", "", err
	}

	return
}
