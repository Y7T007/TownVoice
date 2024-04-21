package auth

import (
	"encoding/json"
	"github.com/ipfs/go-ipfs-api"
	"io/ioutil"
	"log"
	"strings"
)

type Client struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Email        string                 `json:"email"`
	Password     string                 `json:"password"`
	City         string                 `json:"city"`
	Address      string                 `json:"address"`
	Gender       string                 `json:"gender"`
	PhoneNumber  string                 `json:"phone_number"`
	OtherDetails map[string]interface{} `json:"other_details"`
	CID          string                 `json:"cid"` // Add a CID field to store the IPFS CID
}

func SaveClient(user *Client) error {
	// Connect to the ipfs daemon
	sh := shell.NewShell("localhost:5001")

	// Convert the client to JSON
	clientJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Add the client to IPFS
	cid, err := sh.Add(strings.NewReader(string(clientJson)))
	if err != nil {
		return err
	}

	// Store the CID in the user object
	user.CID = cid

	log.Printf("Addiiiing client %s to IPFS with CID %s", user.Email, cid)

	return nil
}

func AuthenticateClient(user *Client) bool {
	// Connect to the ipfs daemon
	sh := shell.NewShell("localhost:5001")

	// Get the client from IPFS
	readCloser, err := sh.Cat(user.CID)
	if err != nil {
		log.Printf("Error retrieving client from IPFS: %v", err)
		return false
	}
	defer readCloser.Close()

	// Read the data from the io.ReadCloser into a []byte
	clientJson, err := ioutil.ReadAll(readCloser)
	if err != nil {
		log.Printf("Error reading client data: %v", err)
		return false
	}

	// Convert the JSON to a client
	var client Client
	err = json.Unmarshal(clientJson, &client)
	if err != nil {
		log.Printf("Error unmarshalling client JSON: %v", err)
		return false
	}

	log.Printf("Retrieved client %s from IPFS with CID %s", client.Email, user.CID)

	// Compare the retrieved client with the input client
	if client.Email == user.Email && client.Password == user.Password {
		return true
	}

	return false
}
