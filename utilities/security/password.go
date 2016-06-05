package security

import (
	"crypto/rand"
	"encoding/base64"

	// Import lib for the v1.08
	"gopkg.in/hlandau/passlib.v1"
)

// PasswordHash is a structure for
// returning the generated results from the
// password
type PasswordHash struct {
	Hash string
	Salt string
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

// GeneratePassword is used to create a secure
// hash with a random salt for the user
func GeneratePassword(password string) (PasswordHash, error) {
	// Generate additional salt for password
	salt, err := GenerateRandomString(32)
	if err != nil {
		return PasswordHash{}, err
	}
	finalPass := password + salt
	hash, hashErr := passlib.Hash(finalPass)
	if hashErr != nil {
		return PasswordHash{}, hashErr
	}
	return PasswordHash{Hash: hash, Salt: salt}, nil
}

// ValidatePassword is a function for validating the user
// password using the stored hash and the user input
func ValidatePassword(hash string, salt string, password string) bool {
	finalPass := password + salt
	_, err := passlib.Verify(finalPass, hash)
	if err != nil {
		return false
	}
	return true
}
