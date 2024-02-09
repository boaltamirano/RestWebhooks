package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

// Contructor de Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0), //Creamos un nuevo slice para clients de longitud 0
		register:   make(chan *Client), // Creamos un canal para register
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

// Definimos la ruta que va a manejar los websockets
// funcion deltro del Hub como un metodo de nombre HandlerWebSocket
func (hub *Hub) HandlerWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error en HandlerWebSocket -> ", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	client := NewClient(hub, socket)
	hub.register <- client // al hub le registramos el cliente

	// Activamos go routina que se encarge de escribir los mensajes al web socket
	go client.Write()
}
