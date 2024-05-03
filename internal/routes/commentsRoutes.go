package routes

import (
	"TownVoice/internal/controllers"
	"TownVoice/utils"
	"net/http"
)

func CommentRoutes(mux *http.ServeMux) {
	mux.Handle("/comments/add-comment/{entityId}", utils.CorsMiddleware((http.HandlerFunc(controllers.AddComment))))
	mux.Handle("/comments/get-comments-by-entity/{entityID}", utils.CorsMiddleware(utils.Middleware(http.HandlerFunc(controllers.GetCommentsByEntity))))
	mux.Handle("/comments/get-comments-by-user/{userID}", utils.CorsMiddleware(utils.Middleware(http.HandlerFunc(controllers.GetCommentsByUser))))
}
