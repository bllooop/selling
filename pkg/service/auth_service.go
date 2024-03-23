package service

import (
	"selling"
	"selling/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user selling.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) ParseToken(accesstok string) (string, error) {
	return "", nil
}

func (s *AuthService) CreateToken(username, password string) (string, error) {
	return "", nil
}
