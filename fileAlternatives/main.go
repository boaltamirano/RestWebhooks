// ALTERNATIVE

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/RestWebkooks/handlers"
	"github.com/RestWebkooks/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseURL := os.Getenv("DATABASE_URL")

	config := &server.Config{
		Port:        ":" + port,
		JWTSecret:   jwtSecret,
		DatabaseUrl: databaseURL,
	}

	httpServer, err := server.NewHTTPServer(context.Background(), config)
	if err != nil {
		log.Fatalf("Error creating HTTP server: %v\n", err)
	}

	bindRoutes(httpServer)

	httpServer.Start()
}

func bindRoutes(s server.Server) {
	httpServer := s.(*server.HTTPServer)
	router := httpServer.Router // Cambiar router por Router

	router.HandleFunc("/", handlers.HomeHandler()).Methods(http.MethodGet)
}
