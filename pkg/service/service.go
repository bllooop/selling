package service

import (
	"selling"
	"selling/pkg/repository"
)

type Authorization interface {
	CreateUser(user selling.User) (int, error)
	CreateToken(username, password string) (string, error)
	ParseToken(accesstok string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
