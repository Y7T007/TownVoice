// internal/auth/controller/login_client.go
package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"net/http"
	"os"
	"time"
)

func GenerateJWT(user *auth.Client) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/client-auth/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse and decode the request body into a new 'Client' type instance
		user := &auth.Client{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Authenticate the user
		isValidClient := auth.AuthenticateClient(user)
		if !isValidClient {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate JWT
		tokenString, err := GenerateJWT(user)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Respond to the client
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}
