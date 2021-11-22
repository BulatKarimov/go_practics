package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"

	"practics/internal/handler"
	"practics/internal/config"
)

func main() {
	server_config := config.Server{}
	server_config_boot_err := cleanenv.ReadConfig("config.yml", &server_config)

	if server_config_boot_err != nil {
		log.Fatal(server_config_boot_err)
	}

	handler := handler.NewHandler()
	router := chi.NewRouter()

	router.Get("/path", handler.Res)

	log.Print("Starting Server on host: " + server_config.Host + " with port " + server_config.Port)
	listener_err := http.ListenAndServe(":" + server_config.Port, router)

	if listener_err != nil {
		log.Fatal(listener_err)
	}
	log.Fatal(listener_err)
}
