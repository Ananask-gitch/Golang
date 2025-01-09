package handlers

import (
	"Golang/pkg/models"
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
	price := r.FormValue("price")
	pricei, e := strconv.Atoi(price)
	if e == nil {
		log.Println("Error")
	}
	advertisement := models.Advertisement{Name: name, Comment: comment, Photo: photo, Price: uint(pricei)}
	result := h.DB.Create(&advertisement)
	if result.Error != nil {
		log.Println(result.Error)
	}
	http.Redirect(w, r, "/advertisements", 301)
}
