package controllers

import (
	"net/http"
)

func AddRating(w http.ResponseWriter, r *http.Request) {
	// Similar to AddComment, but with scores instead of a comment
}

func GetRatingsByEntity(w http.ResponseWriter, r *http.Request) {
	// Similar to GetCommentsByEntity, but with ratings instead of comments
}
