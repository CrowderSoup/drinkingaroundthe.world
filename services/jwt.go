package services

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type JwtService struct {
	secret string
}

func NewJwtService() *JwtService {
	secret := viper.GetString("secret")

	return &JwtService{
		secret: secret,
	}
}

func (j *JwtService) CreateLoginToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	return token.SignedString(j.secret)
}

func (j *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secret), nil
	})
}