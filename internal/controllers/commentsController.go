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

	// Get the entity ID and comment from the URL
	entityId := strings.TrimPrefix(r.URL.Path, "/comments/add-comment/")

	// Decode the request body into a map
	var requestData map[string]string
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the comment from the request data
	comment := requestData["comment"]

	// Create a new Comment object
	newComment := models.Comment{
		UserID:   uid,
		EntityID: entityId,
		Content:  comment,
	}

	// Create a new BadWordDetector visitor
	badWordDetector := &models.BadWordDetector{
		BadWords: []string{"fuck", "badword2"},
	}

	// Use the visitor to check the comment for bad words
	newComment.Accept(badWordDetector)

	// If the comment content is empty, it means it contained bad words
	if newComment.Content == "" {
		http.Error(w, "Comment contains sensitive content", http.StatusBadRequest)
		return
	}

	// Log the user's UID, name, entity ID, and comment
	fmt.Printf("User with UID %s and name %s added a comment on entity %s: %s\n", uid, name, entityId, comment)

	// Implement your logic here to actually add the comment
	facade.AddComment(entityId, comment, uid)

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
