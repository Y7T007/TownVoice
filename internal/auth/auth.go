package auth

import (
	"encoding/hex"
	"golang.org/x/crypto/nacl/auth"
)

var SecretKey [32]byte

func InitSecretKey() {
	secretKeyBytes, err := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	if err != nil {
		panic(err)
	}

	copy(SecretKey[:], secretKeyBytes)
}

func AuthenticateMessage(message []byte) *[32]byte {
	return auth.Sum(message, &SecretKey)
}
