package handler

import (
	"fmt"
	"log"
	"net/http"

	"practics/internal/api"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Res(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting localhost:8080/path request")

	var jokeClient api.JokeClient

	result, error := jokeClient.GetJoke()

	if error != nil {
		fmt.Fprint(w, error)
	}else{
		fmt.Fprint(w, result)
	}
}
