package handler

import (
	"selling/pkg/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/api/auth/sign-up", h.signUp)
	m.HandleFunc("/api/auth/sign-in", h.signIn)
	m.HandleFunc("/api/create-selling", h.AuthMiddleware(h.createSellinglist, "create-selling"))
	m.HandleFunc("/api/sellings", h.AuthMiddleware(h.getAllSelling, "sellings"))
	return m
}
