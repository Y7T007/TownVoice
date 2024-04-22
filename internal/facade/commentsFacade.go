package facade

import (
	"TownVoice/internal/models"
	"TownVoice/internal/repositories/databasesRepo"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"
)

func GetCommentsByEntity(entityID string) {

	//	print that its received
}

func GetCommentsByUser(userID string) {

}

func AddComment(entityID string, comment string, uid string) {
	// Print that it's received
	log.Printf("From the facade: User with UID %s added a comment on entity %s: %s\n", uid, entityID, comment)

	// Create a new Comment object
	newComment := models.Comment{
		UserID:    uid,
		EntityID:  entityID,
		Content:   comment,
		Timestamp: time.Now(),
	}

	// Convert the Comment object to JSON
	commentJSON, err := json.Marshal(newComment)
	if err != nil {
		log.Fatalf("Failed to convert comment to JSON: %v", err)
	}

	// Print the JSON of the new comment
	log.Printf("Comment JSON: %s\n", commentJSON)

	// Get a new IPFSRepo
	ipfsRepo := databasesRepo.NewIPFSRepo()

	// Add the comment JSON to IPFS
	cid, err := ipfsRepo.Shell.Add(strings.NewReader(string(commentJSON)))
	if err != nil {
		log.Fatalf("Failed to add comment to IPFS: %v", err)
	}

	// Print the CID of the new IPFS object
	log.Printf("Added comment to IPFS with CID %s\n", cid)

	// Add the cid to Firestore
	firestoreRepo := databasesRepo.NewFirestoreRepo()
	ctx := context.Background()

	// Fetch existing document
	docRef := firestoreRepo.Client.Collection("Comments").Doc(entityID)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch document: %v", err)
	}

	// Extract existing cids from the document
	var existingCids []string
	if docSnap.Exists() {
		data := docSnap.Data()
		if data != nil {
			if cidList, ok := data["cids"].([]interface{}); ok {
				for _, c := range cidList {
					if cid, ok := c.(string); ok {
						existingCids = append(existingCids, cid)
					}
				}
			}
		}
	}

	// Append the new cid to existing cids
	existingCids = append(existingCids, cid)

	// Update the Firestore document with the new cids
	_, err = docRef.Set(ctx, map[string]interface{}{"cids": existingCids}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed to update Firestore document: %v", err)
	}

	// Print the CID of the new IPFS object
	log.Printf("Added CID to Firestore with CID %s\n", cid)
}
