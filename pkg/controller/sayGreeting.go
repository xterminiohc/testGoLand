package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	userdb "sofka.com/mod/pkg/adapter"
)

var tokenUser = []byte("richardsonmaturity")

func Greet(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Token")

	claims, ok := extractClaims(token)

	if ok {
		idCLient := claims["client"].(string)
		language := claims["language"].(string)

		user := userdb.FindUserbyName(idCLient)

		switch language {
		case "es":
			fmt.Fprintf(w, "Hola, "+user.Name)
		case "it":
			fmt.Fprintf(w, "Ciao, "+user.Name)
		case "en":
			fmt.Fprintf(w, "Hello, "+user.Name)
		}
	}

}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(tokenUser)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
