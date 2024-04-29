package controllers

import (
	"TownVoice/internal/repositories/databasesRepo"
	"encoding/json"
	"net/http"
)

type QRRequest struct {
	TransactionID string   `json:"transaction_id"`
	EntityType    string   `json:"entity_type"`
	Elements      []string `json:"elements"`
	Amount        float64  `json:"amount"`
}

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a QRRequest object
	var qrRequest QRRequest
	err := json.NewDecoder(r.Body).Decode(&qrRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Create a new FirestoreRepo
	firestoreRepo := databasesRepo.NewFirestoreRepo()

	// Store the recipient's ID in Firestore
	_, _, err = firestoreRepo.Client.Collection("recipients").Add(r.Context(), map[string]interface{}{
		"id": qrRequest.TransactionID,
	})
	if err != nil {
		http.Error(w, "Failed to store recipient ID", http.StatusInternalServerError)
		return
	}

}
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	// Implement your payment processing logic here

	// For example, you might get the payment details from the request,
	// call a function from your payment service to process the payment,
	// and then write a response back to the client.

	// Don't forget to handle errors appropriately!
}
