package handlers

import (
	"TownVoice/internal/auth"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Render the login form
		http.ServeFile(w, r, "./web/login.html")
	case "POST":
		// Handle the form submission
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Authenticate the user
		mac := auth.AuthenticateMessage([]byte(password))
		fmt.Printf("MAC: %x\n", *mac)

		// For now, just print the username and MAC
		fmt.Fprintf(w, "Username: %s\nMAC: %x\n", username, *mac)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
