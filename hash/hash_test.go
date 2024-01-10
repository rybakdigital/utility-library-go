package hash

import (
	"fmt"
	"slices"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGenerateHash(t *testing.T) {
	sizes := []int{3, 5, 10}

	for _, length := range sizes {
		hash := GenerateHash(length)
		assert.Equal(t, length, len(hash))

		// Making sure random hash has been generated
		secondHash := GenerateHash(length)
		assert.NotEqual(t, hash, secondHash)
	}
}

func TestGenerateBase64(t *testing.T) {
	sizes := []int{3, 5, 10}

	for _, length := range sizes {
		hash := GenerateBase64(length)
		assert.Equal(t, length, len(hash))

		// Making sure random hash has been generated
		secondHash := GenerateBase64(length)
		assert.NotEqual(t, hash, secondHash)
	}
}

func TestGenerateBaseSpecial(t *testing.T) {
	sizes := []int{3, 5, 10}

	for _, length := range sizes {
		hash := GenerateBaseSpecial(length)
		assert.Equal(t, length, len(hash))

		// Making sure random hash has been generated
		secondHash := GenerateBaseSpecial(length)
		assert.NotEqual(t, hash, secondHash)
	}
}

func TestHash16(t *testing.T) {
	assert.Equal(t, len(Hash16()), 16)
}

func TestHash32(t *testing.T) {
	assert.Equal(t, len(Hash32()), 32)
}

func TestGetHashCharBase62(t *testing.T) {
	// Let's check the scenario number of times, to assure better randomness
	for i := 0; i < 100; i++ {
		// Generate random Base62 character
		char := getHashChar("")

		if !slices.Contains([]byte(base62), char) {
			t.Errorf("Failed to assert that character %s is present in base %s", string(char), base62)
		}
	}
}

func TestGetHashCharBase64(t *testing.T) {
	// Let's check the scenario number of times, to assure better randomness
	for i := 0; i < 100; i++ {
		// Generate random base64 character
		char := getHashChar("base64")

		if !slices.Contains([]byte(base62+base64), char) {
			t.Errorf("Failed to assert that character %s is present in base %s", string(char), base62+base64)
		}
	}
}

func TestGetHashCharBaseSpecial(t *testing.T) {
	// Let's check the scenario number of times, to assure better randomness
	for i := 0; i < 100; i++ {
		// Generate random baseSpecial character
		char := getHashChar("baseSpecial")

		if !slices.Contains([]byte(base62+base64+baseSpecial), char) {
			t.Errorf("Failed to assert that character %s is present in base %s", string(char), base62+base64+baseSpecial)
		}
	}
}

func TestAbleToGenerateBaseSpecialChar(t *testing.T) {
	// Making sure the attempt number is rather large
	attempts := 10000

	for i := 1; i <= attempts; i++ {
		// Generate random baseSpecial character
		char := getHashChar("baseSpecial")

		// If random character is a special one break out and finish test
		if slices.Contains([]byte(baseSpecial), char) {
			break // break here
		}

		// We have tried many times and failed to generate special character
		if i == attempts {
			t.Errorf("Failed to generate character from baseSpecial, tried %d times", i)
		}
	}
}

func BenchmarkGenerateHash(b *testing.B) {
	lengths := []int{3, 5, 10, 15, 20}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("benchmark-length-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateHash(length)
			}
		})
	}
}
