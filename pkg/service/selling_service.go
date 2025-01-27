package service

import (
	"errors"
	"selling"
	"selling/pkg/repository"
	"strings"
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
	if strings.ToLower(order) != "title" || strings.ToLower(order) != "price" || strings.ToLower(order) != "date" {
		return nil, errors.New("order type is incorrect, choose either title, price, date")
	}
	if strings.ToLower(sortby) != "asc" || strings.ToLower(sortby) != "desc" {
		return nil, errors.New("sort type is incorrect, choose either asc or desc")
	}
	return s.repo.ListSellings(userId, order, sortby, page)
}
