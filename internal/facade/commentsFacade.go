package facade

import (
	"TownVoice/internal/models"
	"TownVoice/internal/repositories/databasesRepo"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func GetCommentsByEntity(entityID string) ([]models.Comment, error) {
	// Get a new FirestoreRepo
	firestoreRepo := databasesRepo.NewFirestoreRepo()
	ctx := context.Background()

	// Fetch existing document
	docRef := firestoreRepo.Client.Collection("Comments").Doc(entityID)
	docSnap, err := docRef.Get(ctx)
	if err != nil {
		log.Printf("Failed to fetch document: %v", err)
		return nil, fmt.Errorf("failed to fetch document: %v", err)
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

	// Get a new IPFSRepo
	ipfsRepo := databasesRepo.NewIPFSRepo()

	// Fetch each comment from IPFS using the cids
	var comments []models.Comment
	for _, cid := range existingCids {
		commentReader, err := ipfsRepo.Shell.Cat(cid)
		if err != nil {
			log.Printf("Failed to fetch comment from IPFS: %v", err)
			continue
		}

		// Read the data from the io.ReadCloser into a byte slice
		commentJSON, err := ioutil.ReadAll(commentReader)
		if err != nil {
			log.Printf("Failed to read comment data: %v", err)
			continue
		}

		// Convert the comment JSON to a Comment object
		var comment models.Comment
		err = json.Unmarshal(commentJSON, &comment)
		if err != nil {
			log.Printf("Failed to convert comment JSON to Comment object: %v", err)
			continue
		}

		comments = append(comments, comment)
	}

	return comments, nil
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
