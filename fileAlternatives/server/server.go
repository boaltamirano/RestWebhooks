// ALTERNATIVE

package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Start()
}

type HTTPServer struct {
	Router *mux.Router // Cambiamos a Router con R mayúscula para exportar el campo
	config *Config
}

func NewHTTPServer(ctx context.Context, config *Config) (*HTTPServer, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	return &HTTPServer{
		Router: mux.NewRouter(), // Cambiamos el nombre del campo a Router
		config: config,
	}, nil
}

func (s *HTTPServer) Start() {
	log.Println("Starting server on port", s.config.Port)
	http.ListenAndServe(s.config.Port, s.Router) // Cambiamos s.router a s.Router
}
