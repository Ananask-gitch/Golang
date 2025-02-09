package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerAdd(w http.ResponseWriter, r *http.Request) {
	var res struct {
		Name      string
		Comment   string
		PhotoMain []models.Photo
		Price     uint
	}
	advertisement := models.Advertisement{}

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	Name := res.Name
	Comment := res.Comment
	Price := res.Price
	Photo := res.PhotoMain

	advertisement.Name = Name
	advertisement.Comment = Comment
	advertisement.Photos = Photo
	advertisement.Price = Price

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
