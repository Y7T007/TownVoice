package controllers

import (
	"TownVoice/internal/repositories/databasesRepo"
	"encoding/json"
	"log"
	"net/http"
)

type QRRequest struct {
	TransactionID string   `json:"transaction_id"`
	EntityID      string   `json:"entity_id"`
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

	// Log the received JSON
	log.Printf("Received JSON: %+v\n", qrRequest)

	// Create a new FirestoreRepo
	firestoreRepo := databasesRepo.NewFirestoreRepo()

	// Store the QRRequest object in Firestore
	_, _, err = firestoreRepo.Client.Collection("transactions").Add(r.Context(), qrRequest)
	if err != nil {
		http.Error(w, "Failed to store transaction", http.StatusInternalServerError)
		return
	}
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {

}
