package main

import (
	"fmt"
	"log"
	"net/http"

	"alvesrafa.dev/study/configs"
	"alvesrafa.dev/study/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

	if err != nil {
		log.Printf("Running on port: %d", configs.GetServerPort())
	}
}
