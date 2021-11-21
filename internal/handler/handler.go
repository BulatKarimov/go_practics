package handler

import (
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Res(w http.ResponseWriter, r *http.Request) {
	log.Print("Starting localhost:8080/path request")

	fmt.Fprint(w, "Handler response")
}
