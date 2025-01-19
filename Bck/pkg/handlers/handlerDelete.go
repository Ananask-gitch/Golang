package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) HandlerDelete(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	advertisement := models.Advertisement{}
	result := h.DB.First(&advertisement, "id = ?", id).Delete(&advertisement)
	if result.Error != nil {
		log.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("объявление не найдено")
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(advertisement)
	}
}
