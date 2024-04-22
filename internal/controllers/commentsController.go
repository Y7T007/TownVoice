package controllers

import (
	"TownVoice/utils"
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
	comment := r.URL.RawQuery

	// Log the user's UID, name, entity ID, and comment
	fmt.Printf("User with UID %s and name %s added a comment on entity %s: %s\n", uid, name, entityId, comment)

	// Implement your logic here to actually add the comment
}

func GetCommentsByEntity(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetCommentsByUser(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}
