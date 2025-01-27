package service

import (
	"selling"
	"selling/pkg/repository"
)

type SellingService struct {
	repo repository.Selling
}

func NewSellingService(repo repository.Selling) *SellingService {
	return &SellingService{repo: repo}
}

func (s *SellingService) CreateSelling(userId int, list selling.SellingList) (selling.SellingList, error) {
	return s.repo.CreateSelling(userId, list)
}

func (s *SellingService) ListSellings(userId int, order, sortby string, page int) (map[string]interface{}, error) {
	return s.repo.ListSellings(userId, order, sortby, page)
}
