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
	return m
}
