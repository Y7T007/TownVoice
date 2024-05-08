package databasesRepo

import (
	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type IPFSRepo struct {
	Shell *ipfsapi.Shell
}

func NewIPFSRepo() *IPFSRepo {
	shell := ipfsapi.NewShell("https://my-ipfs-service-e6khoe6iuq-ew.a.run.app/")

	return &IPFSRepo{
		Shell: shell,
	}
}
