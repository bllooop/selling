package service

import (
	"errors"
	"selling"
	"selling/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user selling.User) (selling.User, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) ParseToken(accesstok string) (int, error) {
	token, err := jwt.ParseWithClaims(accesstok, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthService) CreateToken(username, password string) (string, error) {
	user, err := s.repo.SignUser(username, password)
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
