package databasesRepo

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type IPFSRepo struct {
	Shell *ipfsapi.Shell
}

func NewIPFSRepo() *IPFSRepo {
	shell := ipfsapi.NewShell("localhost:5001")

	return &IPFSRepo{
		Shell: shell,
	}
}
