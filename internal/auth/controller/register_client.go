package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Register struct for registration credentials
type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler handles registration requests
func RegisterClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/register", http.StatusSeeOther) // Redirect to register page
		return
	} else if r.Method == "POST" {
		// Parse registration credentials
		var register Register
		err := json.NewDecoder(r.Body).Decode(&register)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate email and password (add your validation logic here)

		// Register user
		if err := registerUser(register.Email, register.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Registration successful, display success message
		w.Write([]byte("Account created successfully"))
	}
}

// registerUser registers a new user (implementation using a separate service)
func registerUser(email, password string) error {
	// Replace with your service logic for user registration using a suitable library
	// (avoid directly using Firebase in the controller)
	err := registerUserService(email, password)
	if err != nil {
		return errors.New("error registering user")
	}

	return nil
}

// registerUserService registers a new user in a separate service (implementation omitted)
func registerUserService(email, password string) error {
	// Implement user registration using a library like `firebase.google.com/go/v4/auth`
	// outside the controller for security best practices

	return errors.New("internal server error") // Placeholder error for registration service
}
