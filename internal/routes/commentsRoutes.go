package routes

import (
	"TownVoice/internal/controllers"
	"net/http"
)

func CommentRoutes(mux *http.ServeMux) {
	mux.Handle("/comments/add-comment/:entityId/:comment", Middleware(http.HandlerFunc(controllers.AddComment)))
	mux.Handle("/comments/get-comments-by-entity/:entityID", Middleware(http.HandlerFunc(controllers.GetCommentsByEntity)))
	mux.Handle("/comments/get-comments-by-user/:userID", Middleware(http.HandlerFunc(controllers.GetCommentsByUser)))
}
