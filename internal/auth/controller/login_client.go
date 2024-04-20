// internal/auth/controller/login_client.go
package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"html/template"
	"net/http"
)

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
		//isValidClient := auth.AuthenticateClient(user)
		//if !isValidClient {
		//	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		//	return
		//}

		// Respond to the client
		w.Write([]byte("Logged in successfully"))
	}
}
