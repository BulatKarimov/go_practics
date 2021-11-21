package main

import (
	"log"
	"net/http"

	"practics/internal/handler"

	"github.com/go-chi/chi/v5"
)

func main() {
	handler := handler.NewHandler()
	router := chi.NewRouter()

	router.Get("/path", handler.Res)

	err := http.ListenAndServe(":8000", router)
	log.Fatal(err)
}
