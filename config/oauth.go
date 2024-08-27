package config

type Oauth struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
}

var GoogleOauthConfig *Oauth
var FacebookOauthConfig *Oauth
