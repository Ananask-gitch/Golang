package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerAdd(w http.ResponseWriter, r *http.Request) {

	advertisement := models.Advertisement{}

	err := json.NewDecoder(r.Body).Decode(&advertisement)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := h.DB.Preload("Photos").Create(&advertisement)
	if result.Error != nil {
		log.Println(result.Error)
		http.Error(w, "Database error", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(advertisement)

}
