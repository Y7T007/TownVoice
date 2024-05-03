package controllers

import (
	"TownVoice/internal/facade"
	"TownVoice/internal/models"
	"TownVoice/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	// Get the entity ID from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/comments/add-comment/")

	// Decode the request body into a map
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the comment and transactionID from the request data
	comment := requestData["comment"].(string)
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

	// Create a new Comment object
	newComment := models.Comment{
		UserID:   transactionID, // use transactionID as UserID
		EntityID: entityId,
		Content:  comment,
	}

	// Create a new BadWordDetector visitor
	badWordDetector := models.NewBadWordDetector()

	// Use the visitor to check the comment for bad words
	newComment.Accept(badWordDetector)

	// If the comment content is empty, it means it contained bad words
	if newComment.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "Comment contains sensitive content",
		})
		return
	}

	// Call the AddComment function from the commentsFacade package
	facade.AddComment(entityId, comment, transactionID) // use transactionID as UserID

	// After the comment is added successfully, write a success status and message back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Comment has been added successfully",
	})
}

func GetCommentsByEntity(w http.ResponseWriter, r *http.Request) {
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
	name := token.Claims["name"]

	// Get the entity ID from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/comments/get-comments-by-entity/")

	// Log the user's UID, name, and entity ID
	fmt.Printf("User with UID %s and name %s requested comments on entity %s\n", uid, name, entityId)

	// Call the GetCommentsByEntity function from the commentsFacade package
	comments, err := facade.GetCommentsByEntity(entityId)
	if err != nil {
		http.Error(w, "No Comment was found for this entity : "+entityId, http.StatusInternalServerError)
		return
	}

	// Write the comments back to the client
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func GetCommentsByUser(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}
