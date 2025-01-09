package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var advertisement []models.Advertisement

	if result := h.DB.First(&advertisement, " ID = ?", id); result.Error != nil {
		log.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(advertisement)
}
