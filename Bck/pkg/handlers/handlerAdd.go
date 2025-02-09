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
	}

	log.Println(advertisement)
	result := h.DB.Preload("Photos").Create(&advertisement)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
