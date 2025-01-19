package handlers

import (
	"Golang/pkg/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h handler) HandlerCreate(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	name := r.FormValue("name")
	comment := r.FormValue("comment")
	photo := r.FormValue("photo")
	photosecond := r.FormValue("photosecond")
	photosecond2 := r.FormValue("photosecond2")
	price := r.FormValue("price")
	pricei, err := strconv.Atoi(price)
	if err != nil {
		log.Println(err)
	}
	if name == "" || photo == "" || comment == "" || price == "" || pricei <= 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Имя, Фото, Комментарий, Цена: имеют пустое значение или некорректное значение цены")
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	} else {
		advertisement := models.Advertisement{
			Name:    name,
			Comment: comment,
			Photos: []models.Photo{{
				PhotoMain:    photo,
				PhotoSecond:  photosecond,
				PhotoSecond2: photosecond2}},
			Price: uint(pricei)}

		result := h.DB.Create(&advertisement)
		if result.Error != nil {
			log.Println(result.Error)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
		}
		h.DB.Save(&advertisement)

	}
}
