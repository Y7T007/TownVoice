package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_PATH"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}

	return token, nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			http.Error(w, "Authorization header not provided", http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authorizationHeader, "Bearer ")
		token, err := VerifyIDToken(r.Context(), idToken)
		if err != nil {
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}

		// You can access the user's Firebase UID as follows:
		uid := token.UID
		log.Printf("User with UID %s authenticated", uid, "the name is : ", token.Claims["name"])
		// Now you can use this uid to perform requests and queries.

		next.ServeHTTP(w, r)
	})
}
