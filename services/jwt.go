package services

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type JwtService struct {
	secret []byte
}

func NewJwtService() *JwtService {
	secret := viper.GetString("secret")

	return &JwtService{
		secret: []byte(secret),
	}
}

func (j *JwtService) CreateLoginToken(claims jwt.MapClaims) (string, error) {
	// TODO:This token should expire after 1 hour (email can be slow) to be more secure
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

func (j *JwtService) ValidateToken(tokenString string, claims jwt.MapClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.secret, nil
	})
}
