// cmd/TownVoice/main.go
package main

import (
	"TownVoice/internal/ipfs"
	"TownVoice/internal/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	// Start the server regardless of whether the file was successfully added to IPFS
	go server.Start(os.Getenv("APP_PORT"))

	// Keep the main function from exiting
	select {}
}
