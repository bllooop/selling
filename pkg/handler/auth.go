package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	selling "selling"
)

func JSONStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	var input selling.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Password == "" || input.Username == "" {
		clientErr(w, http.StatusBadRequest, "invalid input body")
		return
	}
	list, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(list)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	var input signInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Password == "" || input.Username == "" {
		clientErr(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.CreateToken(input.Username, input.Password)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})
	r.Header.Set("Authorization", "Bearer "+token)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
