package server

// Esta clase explica como codificar el servidor, las siguientes notas te pueden ayudar:

// Estructuras:
//    Config: tiene las caracteristicas del servidor. El puerto en el que se va ejecutar, la clave secreta para generar tokens y la conexion a base de datos.
//    broker: Nos ayuda a tener varias instancias de servidor corriendo. Esta estructura a su vez tiene la estructura Config y el metodo Config, para ser de tipo Server.

// Interface:
//    server: esta interface implementa el modelo de datos o estructura de config.

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/RestWebkooks/database"
	"github.com/RestWebkooks/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	websocket "github.com/RestWebkooks/websocket"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

// Para tener Server tendremos un metodo Config() que retorne algo de tipo *config(type Config struct)
type Server interface {
	Config() *Config
	Hub() *websocket.Hub // implementamos una nueva funcion
}

// Broker que se encargara de manejar este server
type Broker struct {
	config *Config
	router *mux.Router // Ruteador para definir las rutas del server
	hub    *websocket.Hub
}

// Metodo para que el struct Broker satisfaga la interface:
// //	Un metodo llamado Config() que retorna una configuracion "*Config"
func (b *Broker) Config() *Config {
	return b.config // retornamos config de Broker para que se comporte como un tipo server "config *Config"
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

// Definir el constructor para nuestro struc
// Creamos una funcion que recibe 2 parametro ctx y config
// Retornamos dos valores "*Broker, error"
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	// Creamos una instancia del broker con la config y router
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}

	return broker, nil
}

// funcion o metodo para levantar el servidor
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter() //instanciamos un nuevo router usando la libreria de mux
	binder(b, b.router)
	handler := cors.AllowAll().Handler(b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	go b.hub.Run() // Activamos el hub.Run para que corra en una subrutina
	repository.SetRepository(repo)
	log.Println("starting server on port", b.config.Port)

	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
