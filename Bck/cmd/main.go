package main

import (
	"Golang/pkg/handlers"
	"Golang/pkg/storage"
	"Golang/pkg/web"
	"log"
	"net/http"
)

func main() {
	mux := web.MakeMux()
	DB := storage.Init()
	h := handlers.New(DB)

	go mux.HandleFunc("GET /advertisements", h.HandlerGetAll)
	go mux.HandleFunc("GET /advertisements/{id}", h.HandlerGet)
	go mux.HandleFunc("DELETE /advertisements/{id}", h.HandlerDelete)
	go mux.HandleFunc("POST /advertisements", h.HandlerAdd)

	log.Println("Listening for requests")
	http.ListenAndServe(":8080", mux)
}
