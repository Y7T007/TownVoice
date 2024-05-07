// ipfs/ipfs.go
package ipfs

import (
	"fmt"
	"github.com/ipfs/go-ipfs-api"
	"strings"
)

// Shell is an IPFS shell
var Shell *shell.Shell

func init() {
	// Create a new IPFS shell pointing to the API address
	Shell = shell.NewShell("/ip4/127.0.0.1/tcp/5002")
}

// AddFile adds a file to IPFS and returns its CID
func AddFile(content string) (string, error) {
	cid, err := Shell.Add(strings.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("error adding file to IPFS: %w", err)
	}
	return cid, nil
}
