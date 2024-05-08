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
	Shell = shell.NewShell("https://my-ipfs-service-e6khoe6iuq-ew.a.run.app/")
}

// AddFile adds a file to IPFS and returns its CID
func AddFile(content string) (string, error) {
	cid, err := Shell.Add(strings.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("error adding file to IPFS: %w", err)
	}
	return cid, nil
}
