package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/db"
	"github.com/tutorials/go/crud/pkg/handlers"
)

func main() {
	log.Println("Hello, World!")
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/v2/modify-points", h.ModifyPoints).Methods(http.MethodPost)
	router.HandleFunc("/v2/points", h.GetAllPoints).Methods(http.MethodGet)
	router.HandleFunc("/v2/points/{id}", h.GetPoint).Methods(http.MethodGet)
	router.HandleFunc("/v2/health", h.Health).Methods(http.MethodGet)
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
