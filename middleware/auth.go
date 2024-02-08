package middleware

import (
	"net/http"
	"strings"

	"github.com/RestWebkooks/models"
	"github.com/RestWebkooks/server"
	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthNiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !shouldCheckToken(r.URL.Path) { // Revisa si la ruta de la peticion debe ser revisada y bloqueada con un token
				next.ServeHTTP(w, r) // Si no debe ser revisada envia al siguiente proceso
				return
			}

			// Siguiente paso para las rutas protegidas
			tokenString := strings.TrimSpace(r.Header.Get("Authorization")) // tokenString -> el token que se envia en la peticion
			_, err := jwt.ParseWithClaims(tokenString, *&models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
