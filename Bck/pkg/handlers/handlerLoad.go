package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerLoad(w http.ResponseWriter, r *http.Request) {

	// Открываем CORS для всех источников
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var advertisements []models.Advertisement

	result := h.DB.Preload("Photos").Find(&advertisements)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(advertisements)
}
