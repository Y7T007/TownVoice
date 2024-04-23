package controllers

import (
	"TownVoice/internal/facade"
	"TownVoice/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// thats working
func AddRating(w http.ResponseWriter, r *http.Request) {
	// Get the user's JWT from the Authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		http.Error(w, "Authorization header not provided", http.StatusUnauthorized)
		return
	}

	// Trim the "Bearer " prefix from the JWT
	idToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Get the user's Firebase UID and other claims from the token
	token, err := utils.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}
	uid := token.UID

	// Get the entity ID from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/ratings/add-rating/")

	// Decode the request body into a map
	var requestData map[string]map[string]float64
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the scores from the request data
	scores := requestData["scores"]

	// Call the AddRating function from the ratingsFacade package
	facade.AddRating(entityId, scores, uid)

	// After the rating is added successfully, write a success status and message back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Rating has been added successfully",
	})
}

// Also workiiiinnnnnggggg
func GetRatingsByEntity(w http.ResponseWriter, r *http.Request) {
	// Get the user's JWT from the Authorization header
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		http.Error(w, "Authorization header not provided", http.StatusUnauthorized)
		return
	}

	// Trim the "Bearer " prefix from the JWT
	idToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

	// Get the user's Firebase UID and other claims from the token
	token, err := utils.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}
	uid := token.UID
	fmt.Printf("The uid is :", uid)
	// Get the entity ID from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/ratings/get-ratings-by-entity/")

	// Call the GetRatingsByEntity function from the ratingsFacade package
	ratings, err := facade.GetRatingsByEntity(entityId)
	if err != nil {
		http.Error(w, "No Rating was found for this entity : "+entityId, http.StatusInternalServerError)
		return
	}

	// Write the ratings back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ratings)
}
