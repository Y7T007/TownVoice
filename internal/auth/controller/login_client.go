package controller

import (
	"TownVoice/internal/models/auth" // Import the correct package
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"html/template"
	"net/http"
)

func LoginClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/client-auth/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Parse and decode the request body into a new 'Client' type instance
		client := &auth.Client{} // Use the correct Client type
		err := json.NewDecoder(r.Body).Decode(client)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Initialize Firebase
		opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			http.Error(w, "Error initializing Firebase", http.StatusInternalServerError)
			return
		}

		// Get a reference to the auth service
		auth, err := app.Auth(context.Background())
		if err != nil {
			http.Error(w, "Error getting Auth client", http.StatusInternalServerError)
			return
		}

		// Verify the ID token
		_, err = auth.VerifyIDToken(context.Background(), client.Password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Respond to the client
		w.Write([]byte("Logged in successfully"))
	}
}
