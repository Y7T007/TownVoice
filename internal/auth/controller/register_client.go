// internal/auth/controller/register_client.go
package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"html/template"
	"net/http"
)

func RegisterClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/client-auth/register.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse and decode the request body into a new 'Client' type instance
		user := &auth.Client{}
		err := json.NewDecoder(r.Body).Decode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Save the user to the database (or in this case, IPFS)
		err = auth.SaveClient(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond to the client
		w.Write([]byte("Registered successfully"))
	}
}
