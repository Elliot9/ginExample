package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

type UserInfo struct {
	Sub     string `json:"sub"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	Agent   Agent  `json:"-"`
}

func (s *service) Callback(agent Agent, state, code string) (userInfo *UserInfo, err error) {
	// todo state check
	config := s.GetOauthConfig(agent)
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := config.Client(context.Background(), token)

	switch agent {
	case agentGoogle:
		res, getErr := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if getErr != nil {
			return nil, getErr
		}
		defer res.Body.Close()

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return nil, readErr
		}

		user := &UserInfo{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			return nil, err
		}

		user.Agent = agent
		return user, nil

	case agentFB:
		res, getErr := client.Get("https://graph.facebook.com/v18.0/me?fields=id,name,email,picture")
		if getErr != nil {
			fmt.Println(getErr)
			return nil, getErr
		}
		defer res.Body.Close()

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return nil, readErr
		}

		user := &UserInfo{}
		err = json.Unmarshal(body, &user)
		if err != nil {
			return nil, err
		}
		user.Agent = agent
		return user, nil
	}

	return nil, fmt.Errorf("agent not found")
}
