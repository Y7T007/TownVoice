package controllers

import (
	"TownVoice/internal/facade"
	"TownVoice/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AddRating(w http.ResponseWriter, r *http.Request) {
	// Get the entity ID from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/ratings/add-rating/")

	// Decode the request body into a map
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the scores and transactionID from the request data
	scoresData := requestData["scores"].(map[string]interface{})
	scores := make(map[string]float64)
	for k, v := range scoresData {
		if value, ok := v.(float64); ok {
			scores[k] = value
		} else {
			http.Error(w, "Invalid score value", http.StatusBadRequest)
			return
		}
	}
	transactionID := requestData["transactionID"].(string)

	// Check if the entityID exists in the "transactions" collection and if it contains the "transaction_ids" field
	exists, err := facade.CheckTransaction(entityId, transactionID)
	if err != nil {
		http.Error(w, "Error checking transaction", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	// Call the AddRating function from the ratingsFacade package
	facade.AddRating(entityId, scores, transactionID)

	// After the rating is added successfully, delete the transaction id
	err = facade.DeleteTransaction(entityId, transactionID)
	if err != nil {
		http.Error(w, "Error deleting transaction", http.StatusInternalServerError)
		return
	}

	// Write a success status and message back to the client
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
