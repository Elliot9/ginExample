package oauth

func (s *service) GetQuery(agent Agent) string {
	config := s.GetOauthConfig(agent)
	return config.AuthCodeURL("csrf")
}
