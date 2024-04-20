// internal/auth/controller/login_client.go
package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"net/http"
)

func LoginClient(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new 'User' type instance
	user := &auth.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate the user
	isValidUser := auth.AuthenticateUser(user)
	if !isValidUser {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Respond to the client
	w.Write([]byte("Logged in successfully"))
}
