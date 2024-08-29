package jwt

import (
	"encoding/hex"
	"github/elliot9/ginExample/config"
	"time"

	"crypto/rand"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(subject string, claims map[string]any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    subject,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
		"claims": claims,
	})

	return token.SignedString([]byte(config.AppSetting.JwtSecret))
}

func GenerateRefreshToken() string {
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)
	return hex.EncodeToString(tokenBytes)
}

func VerifyTokenAndGetClaims(token string) (*jwt.Claims, error) {
	certificateToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppSetting.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return &certificateToken.Claims, nil
}
