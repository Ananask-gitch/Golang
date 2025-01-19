package handlers

import (
	"log"
	"net/http"
)

func (h handler) HandlerGetOne(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var results []struct {
		Name string
	}

	result := h.DB.Table("advertisements").
		Joins("left join photos on photos.advertisement_id = 1").
		Where("id = ?", 1).
		Scan(&results)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
