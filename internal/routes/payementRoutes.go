package routes

import (
	"TownVoice/internal/controllers" // Assuming your payment controllers are in this package
	"TownVoice/utils"
	"net/http"
)

func PaymentRoutes(mux *http.ServeMux) {
	mux.Handle("/api/service-payement", utils.CorsMiddleware(http.HandlerFunc(controllers.ProcessPayment)))
	mux.Handle("/api/Generate-QR-Code", utils.CorsMiddleware(http.HandlerFunc(controllers.GenerateQRCode)))
}
