package hash

import (
	"math/rand"
	"time"
)

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base64 = "-_"
const baseSpecial = "!?@$#%^&*()+="

// GenerateHash generates random hash from base62 (alphanumeric) with a given length
func GenerateHash(length int) string {
	return generate("base62", length)
}

// GenerateBase64 generates random hash from base64 (alphanumeric plus -_) with a given length
func GenerateBase64(length int) string {
	return generate("base64", length)
}

// GenerateHashSpecial generates random hash from baseSpecial (base62, base64 and baseSpecial) with a given length
func GenerateBaseSpecial(length int) string {
	return generate("baseSpecial", length)
}

func generate(base string, length int) string {
	hash := make([]byte, length)

	// Add random characters to hash
	for i := range hash {
		hash[i] = getHashChar(base)
	}

	return string(hash)
}

func getHashChar(baseType string) byte {
	// Add random seed
	rand.NewSource(time.Now().UnixNano())

	// Select base to use
	base := base62
	switch baseType {
	case "base64":
		base += base64
	case "baseSpecial":
		base += base64 + baseSpecial
	}

	return base[rand.Int63()%int64(len(base))]
}

// Hash16 Helper function to quickly generate hash of 16 character length
func Hash16() string {
	return GenerateHash(16)
}

// Hash32 Helper function to quickly generate hash of 32 character length
func Hash32() string {
	return GenerateHash(32)
}
