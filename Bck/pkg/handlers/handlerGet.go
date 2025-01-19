package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerGet(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var results []struct {
		ID           uint
		Name         string
		Comment      string
		PhotoMain    string
		PhotoSecond  string
		PhotoSecond2 string
		Price        uint
	}

	result := h.DB.Table("advertisements").
		Joins("left join photos on photos.advertisement_id = advertisements.id").
		Where("id = ?", id).
		Scan(&results)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&results)
}
