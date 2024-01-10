package hash

import (
	"math/rand"
	"time"
)

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateHash generates random hash from base62 (alphanumeric) with a given length
func GenerateHash(length int) string {
	// Add random seed
	rand.NewSource(time.Now().UnixNano())

	hash := make([]byte, length)

	// Add random characters to hash
	for i := range hash {
		hash[i] = base62[rand.Int63()%int64(len(base62))]
	}

	return string(hash)
}

// Hash16 Helper function to quickly generate hash of 16 character length
func Hash16() string {
	return GenerateHash(16)
}

// Hash32 Helper function to quickly generate hash of 32 character length
func Hash32() string {
	return GenerateHash(32)
}
