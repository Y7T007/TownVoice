package routes

import (
	"TownVoice/internal/controllers"
	"TownVoice/utils"
	"net/http"
)

func RatingRoutes(mux *http.ServeMux) {
	mux.Handle("/ratings/add-rating/{entityId}", utils.CorsMiddleware((utils.Middleware(http.HandlerFunc(controllers.AddRating)))))
	mux.Handle("/ratings/get-ratings-by-entity/{entityID}", utils.CorsMiddleware(utils.Middleware(http.HandlerFunc(controllers.GetRatingsByEntity))))
}
