// cmd/TownVoice/main.go
package main

import (
	"TownVoice/internal/ipfs"
	"TownVoice/internal/server"
	"fmt"
	"log"
)

func main() {
	cid, err := ipfs.AddFile("Hello, IPFS!")
	if err != nil {
		log.Printf("error adding file to IPFS: %v", err) // Log the error and continue
	} else {
		fmt.Println("added file:", cid)
	}

	// Start the server regardless of whether the file was successfully added to IPFS
	go server.Start()

	// Keep the main function from exiting
	select {}
}
