package main

import (
	"TownVoice/internal/ipfs"
	"TownVoice/internal/server"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cid, err := ipfs.AddFile(os.Getenv("IPFS_MESSAGE"))
	if err != nil {
		log.Printf("error adding file to IPFS: %v", err) // Log the error and continue
	} else {
		fmt.Println("added file:", cid)
	}

	srv := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
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
