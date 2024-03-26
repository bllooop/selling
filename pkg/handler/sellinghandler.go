package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"selling"
	"strconv"
)

func (h *Handler) createSellinglist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	retrievedValue := r.Context().Value(idCtx).(int)
	var input selling.SellingList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Title == "" || input.Description == "" || input.Price == 0 || input.PicURL == "" {
		clientErr(w, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Selling.CreateSelling(retrievedValue, input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

func (h *Handler) getAllSelling(w http.ResponseWriter, r *http.Request) {
	order := r.URL.Query().Get("order")
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if r.URL.Path != "/api/movies" {
		notFound(w)
		return
	}
	if err != nil {
		servErr(w, err, err.Error())
	}
	retrievedValue := r.Context().Value(idCtx).(int)
	lists, err := h.services.ListSellings(retrievedValue, order, page)
	if err != nil {
		servErr(w, err, err.Error())
	}
	res, err := JSONStruct(lists)
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}
