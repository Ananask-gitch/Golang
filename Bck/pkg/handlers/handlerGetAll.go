package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerGetAll(w http.ResponseWriter, r *http.Request) {

	// Открываем CORS для всех источников
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var results []struct {
		ID        uint
		Name      string
		Comment   string
		Price     uint
		PhotoMain string
	}

	result := h.DB.Table("advertisements").
		Select("advertisements.id ,advertisements.name, advertisements.comment,advertisements.price , photos.photo_main").
		Joins("left join photos on photos.advertisement_id = advertisements.id").
		Scan(&results)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}
