package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	idCtx               = "userID"
)

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get(authorizationHeader)
		if tokenString == "" {
			clientErr(w, http.StatusUnauthorized, "Authorization token is required")
			return
		}
		headerSplit := strings.Split(tokenString, " ")
		if len(headerSplit) != 2 {
			clientErr(w, http.StatusUnauthorized, "invalid auth header")
			return
		}
		userId, err := h.services.Authorization.ParseToken(headerSplit[1])
		if err != nil {
			clientErr(w, http.StatusUnauthorized, "invalid auth header")
			return
		}
		r = setValue(r, userId)
		next.ServeHTTP(w, r)
	})
}
func setValue(r *http.Request, userId string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), idCtx, userId))
}
