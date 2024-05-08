//go:build heroku
// +build heroku

// main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"TownVoice/internal/ipfs"
	"TownVoice/internal/server"

	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cid, err := ipfs.AddFile("Hello world")
	if err != nil {
		log.Printf("error adding file to IPFS: %v", err) // Log the error and continue
	} else {
		fmt.Println("added file:", cid)
	}

	// Initialize Firebase
	opt := option.WithCredentialsFile("./internal/config/pringles.json")
	_, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server.SetupRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
