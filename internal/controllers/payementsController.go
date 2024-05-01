package controllers

import (
	"TownVoice/internal/repositories/databasesRepo"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type QRRequest struct {
	TransactionID string   `json:"transaction_id"`
	EntityID      string   `json:"entity_id"`
	EntityType    string   `json:"entity_type"`
	Elements      []string `json:"elements"`
	Amount        float64  `json:"amount"`
}

func GenerateQRCode(w http.ResponseWriter, r *http.Request) {
	/// Decode the request body into a QRRequest object
	var qrRequest QRRequest
	err := json.NewDecoder(r.Body).Decode(&qrRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log the received JSON
	log.Printf("Received JSON: %+v\n", qrRequest)

	// Get a new FirestoreRepo
	firestoreRepo := databasesRepo.NewFirestoreRepo()

	// Check if a document with the entityID exists in the "transactions" collection
	docRef := firestoreRepo.Client.Collection("transactions").Doc(qrRequest.EntityID)
	docSnap, err := docRef.Get(context.Background())
	if status.Code(err) == codes.NotFound {
		// If the document does not exist, create a new one with the transactionID in the "transaction_ids" field
		_, err = docRef.Create(context.Background(), map[string]interface{}{"transaction_ids": []string{qrRequest.TransactionID}})
		if err != nil {
			http.Error(w, "Failed to create new document", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, "Error checking document", http.StatusInternalServerError)
		return
	} else {
		// If the document exists, append the transactionID to the "transaction_ids" field
		transactionIds, _ := docSnap.Data()["transaction_ids"].([]interface{})
		transactionIds = append(transactionIds, qrRequest.TransactionID)
		_, err = docRef.Set(context.Background(), map[string]interface{}{"transaction_ids": transactionIds}, firestore.MergeAll)
		if err != nil {
			http.Error(w, "Failed to update document", http.StatusInternalServerError)
			return
		}
	}
	// Log the received JSON
	log.Printf("Received JSON: %+v\n", qrRequest)

	// Open the JSON file
	jsonFile, err := os.Open("internal/config/entities_scores.json")
	if err != nil {
		http.Error(w, "Failed to open JSON file", http.StatusInternalServerError)
		return
	}
	defer jsonFile.Close()

	// Read the file into a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Create a map to hold the JSON data
	var result map[string]interface{}

	// Unmarshal the byte array into the map
	json.Unmarshal([]byte(byteValue), &result)

	// Access the fields of the entity type
	fields := result[qrRequest.EntityType]

	// Convert the fields to JSON
	fieldsJSON, err := json.Marshal(fields)
	if err != nil {
		http.Error(w, "Failed to convert fields to JSON", http.StatusInternalServerError)
		return
	}

	// Write the JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(fieldsJSON)
}

func ProcessPayment(w http.ResponseWriter, r *http.Request) {

}
