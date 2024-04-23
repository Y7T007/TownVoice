package facade

import (
	"TownVoice/internal/models"
	"TownVoice/internal/repositories/databasesRepo"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"strings"
)

func AddRating(entityID string, scores map[string]float64, uid string) {
	// Print that it's received
	log.Printf("From the facade: User with UID %s added a rating on entity %s: %v\n", uid, entityID, scores)

	// Create a new Rating object
	newRating := models.Rating{
		UserID:   uid,
		EntityID: entityID,
		Scores:   scores,
	}

	// Convert the Rating object to JSON
	ratingJSON, err := json.Marshal(newRating)
	if err != nil {
		log.Fatalf("Failed to convert rating to JSON: %v", err)
	}

	// Print the JSON of the new rating
	log.Printf("Rating JSON: %s\n", ratingJSON)

	// Get a new IPFSRepo
	ipfsRepo := databasesRepo.NewIPFSRepo()

	// Add the rating JSON to IPFS
	cid, err := ipfsRepo.Shell.Add(strings.NewReader(string(ratingJSON)))
	if err != nil {
		log.Fatalf("Failed to add rating to IPFS: %v", err)
	}

	// Print the CID of the new IPFS object
	log.Printf("Added rating to IPFS with CID %s\n", cid)

	// Add the cid to Firestore
	firestoreRepo := databasesRepo.NewFirestoreRepo()
	ctx := context.Background()

	// Fetch existing document
	docRef := firestoreRepo.Client.Collection("Ratings").Doc(entityID)
	docSnap, err := docRef.Get(ctx)
	if status.Code(err) == codes.NotFound {
		// If the document does not exist, create a new one
		_, err = docRef.Create(ctx, map[string]interface{}{"cids": []string{cid}})
		if err != nil {
			log.Fatalf("Failed to create new document: %v", err)
		}
	} else if err != nil {
		log.Fatalf("Failed to fetch document: %v", err)
	} else {
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
	}
	// Print the CID of the new IPFS object
	log.Printf("Added CID to Firestore with CID %s\n", cid)
}

func GetRatingsByEntity(entityID string) ([]models.Rating, error) {
	// Get a new FirestoreRepo
	firestoreRepo := databasesRepo.NewFirestoreRepo()
	ctx := context.Background()

	// Fetch existing document
	docRef := firestoreRepo.Client.Collection("Ratings").Doc(entityID)
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

	// Fetch each rating from IPFS using the cids
	var ratings []models.Rating
	for _, cid := range existingCids {
		ratingReader, err := ipfsRepo.Shell.Cat(cid)
		if err != nil {
			log.Printf("Failed to fetch rating from IPFS: %v", err)
			continue
		}

		// Read the data from the io.ReadCloser into a byte slice
		ratingJSON, err := ioutil.ReadAll(ratingReader)
		if err != nil {
			log.Printf("Failed to read rating data: %v", err)
			continue
		}

		// Convert the rating JSON to a Rating object
		var rating models.Rating
		err = json.Unmarshal(ratingJSON, &rating)
		if err != nil {
			log.Printf("Failed to convert rating JSON to Rating object: %v", err)
			continue
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}
