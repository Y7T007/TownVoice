package controller

import (
	"TownVoice/internal/models/auth"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func RegisterClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/client-auth/register.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse and decode the request body into a new 'Client' type instance
		client := &auth.Client{}
		err := json.NewDecoder(r.Body).Decode(client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Log the received form data
		log.Printf("Received form data from the view  : %+v\n", client)

		// Save the client to the IPFS database
		err = auth.SaveClient(client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond to the client
		w.Write([]byte("Registered successfully"))
	}
}
