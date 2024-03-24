package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"selling"
)

func (h *Handler) createSellinglist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	//retrievedValue := "1" // when testing uncomment
	retrievedValue := r.Context().Value(idCtx).(int) // when testing comment
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
