package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) HandlerGetAll(w http.ResponseWriter, r *http.Request) {
	advertisements := []models.Advertisement{}

	result := h.DB.Find(&advertisements)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(advertisements)
}
