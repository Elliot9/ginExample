package oauth

import (
	"github.com/elliot9/gin-example/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (s *service) GetOauthConfig(agent Agent) *oauth2.Config {
	switch agent {
	case agentGoogle:
		return &oauth2.Config{
			ClientID:     config.GoogleOauthConfig.ClientID,
			ClientSecret: config.GoogleOauthConfig.ClientSecret,
			RedirectURL:  "http://localhost:3000/auth/google/callback",
			Scopes:       config.GoogleOauthConfig.Scopes,
			Endpoint:     google.Endpoint,
		}
	case agentFB:
		return &oauth2.Config{
			ClientID:     config.FacebookOauthConfig.ClientID,
			ClientSecret: config.FacebookOauthConfig.ClientSecret,
			RedirectURL:  "http://localhost:3000/auth/facebook/callback",
			Scopes:       config.FacebookOauthConfig.Scopes,
			Endpoint:     google.Endpoint,
		}
	}
	return nil
}
