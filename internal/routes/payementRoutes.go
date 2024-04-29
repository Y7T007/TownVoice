package routes

import (
	"TownVoice/internal/controllers" // Assuming your payment controllers are in this package
	"net/http"
)

func PaymentRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/service-payement", controllers.ProcessPayment)
	mux.HandleFunc("/api/Generate-QR-Code", controllers.GenerateQRCode)
}
