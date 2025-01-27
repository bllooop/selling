package service

import (
	"selling"
	"selling/pkg/repository"
)

type Authorization interface {
	CreateUser(user selling.User) (selling.User, error)
	CreateToken(username, password string) (string, error)
	ParseToken(accesstok string) (int, error)
}

type Selling interface {
	CreateSelling(userId int, list selling.SellingList) (selling.SellingList, error)
	ListSellings(userId int, order, sortby string, page int) (map[string]interface{}, error)
}

type Service struct {
	Authorization
	Selling
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Selling:       NewSellingService(repos.Selling),
	}
}
