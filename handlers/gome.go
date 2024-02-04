package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RestWebkooks/server"
)

type HomeResponse struct {
	Message string `json:"message"` // Especificamos que en go la propuedad se llamara "Message" pero cuando se serealize a json tomara otro nombre "message"
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Conten-Type", "application/json")
		w.WriteHeader((http.StatusOK))
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to the First test to Omar",
			Status:  true,
		})
	}
}
