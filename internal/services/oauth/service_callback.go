package oauth

import (
	"context"
	"fmt"
	"github/elliot9/ginExample/internal/models"
)

func (s *service) Callback(agent Agent, state, code string) (user *models.User, err error) {
	// todo state check
	config := s.GetOauthConfig(agent)
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)
	fmt.Println(client)

	return nil, nil
}
