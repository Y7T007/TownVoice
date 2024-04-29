package controllers

import (
	"TownVoice/internal/repositories/databasesRepo"
	"cloud.google.com/go/firestore"
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

	// Get a reference to a document named after the entity_id
	docRef := firestoreRepo.Client.Collection("transactions").Doc(qrRequest.EntityID)

	// Get the current data in the document (if it exists)
	doc, err := docRef.Get(r.Context())
	if err != nil && !doc.Exists() {
		// If the document doesn't exist, create it with the transaction_id as the first item in an array
		_, err = docRef.Set(r.Context(), map[string][]string{
			"transaction_ids": {qrRequest.TransactionID},
		})
		if err != nil {
			http.Error(w, "Failed to create document", http.StatusInternalServerError)
			return
		}
	} else {
		// If the document exists, append the transaction_id to the existing array
		_, err = docRef.Update(r.Context(), []firestore.Update{
			{
				Path:  "transaction_ids",
				Value: firestore.ArrayUnion(qrRequest.TransactionID),
			},
		})
		if err != nil {
			http.Error(w, "Failed to update document", http.StatusInternalServerError)
			return
		}
	}
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {

}
