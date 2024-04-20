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
		log.Fatalf("error adding file to IPFS: %v", err)
	}
	fmt.Println("added file:", cid)

	server.Start()
}
