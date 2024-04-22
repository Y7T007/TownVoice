package routes

import (
	"TownVoice/internal/controllers"
	"TownVoice/utils"
	"net/http"
)

func CommentRoutes(mux *http.ServeMux) {
	mux.Handle("/comments/add-comment/{entityId}", utils.Middleware(http.HandlerFunc(controllers.AddComment)))
	mux.Handle("/comments/get-comments-by-entity/{entityID}", utils.Middleware(http.HandlerFunc(controllers.GetCommentsByEntity)))
	mux.Handle("/comments/get-comments-by-user/{userID}", utils.Middleware(http.HandlerFunc(controllers.GetCommentsByUser)))
}
