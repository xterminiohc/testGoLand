package auth

import (
	"fmt"
	"net/http"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var tokenUser = []byte("richardsonmaturity")

func GetToken(w http.ResponseWriter, r *http.Request) {
	validTokenSpanish, err := generateJWT("es")
	validTokenEnglish, err := generateJWT("en")
	validTokenItalic, err := generateJWT("it")
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprintf(w, "Token para español:  "+validTokenSpanish)
	fmt.Fprintf(w, "\nEnglish Token:  "+validTokenEnglish)
	fmt.Fprintf(w, "\nGettone per l'italiano:  "+validTokenItalic)
}

func generateJWT(language string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Edwin"
	claims["language"] = language
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(tokenUser)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return tokenUser, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
