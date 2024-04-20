// internal/auth/controller/register_client.go
package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"net/http"
)

func RegisterClient(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new 'User' type instance
	user := &auth.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save the user to the database (or in this case, IPFS)
	err = auth.SaveUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.Write([]byte("Registered successfully"))
}
