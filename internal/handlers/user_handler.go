package handlers

import (
	"TownVoice/internal/auth"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		// Construct the absolute path to the login.html file
		loginFilePath := filepath.Join(cwd, "web", "login.html")

		// Render the login form
		http.ServeFile(w, r, loginFilePath)
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
