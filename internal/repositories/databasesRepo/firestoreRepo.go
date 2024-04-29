package databasesRepo

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreRepo struct {
	Client *firestore.Client
}

func NewFirestoreRepo() *FirestoreRepo {
	ctx := context.Background()

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_PATH"))

	client, err := firestore.NewClient(ctx, "townvoice-485fb", opt)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return &FirestoreRepo{
		Client: client,
	}
}
