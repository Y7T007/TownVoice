package facade

import (
	"TownVoice/internal/models"
	"TownVoice/internal/repositories/databasesRepo"
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
	log.Printf("From the facade : User with UID %s and  added a comment on entity %s: %s\n", uid, entityID, comment)

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
}
