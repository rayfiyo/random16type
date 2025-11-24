package main

import (
	"log"
	"net/http"

	"github.com/rayfiyo/random16type/internal/adapter/httpserver"
	"github.com/rayfiyo/random16type/internal/infrastructure/random"
	"github.com/rayfiyo/random16type/internal/usecase/typecode"
)

func main() {
	randomGenerator := random.New()
	generatorUsecase := typecode.NewGeneratorUsecase(randomGenerator)
	pageHandler := httpserver.NewPageHandler(generatorUsecase)

	mux := http.NewServeMux()
	mux.Handle("/", pageHandler)

	port := ":8080"
	log.Printf("serving random16type at http://localhost%s/", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
