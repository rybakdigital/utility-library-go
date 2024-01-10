package hash

import (
	"fmt"
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

func TestHash16(t *testing.T) {
	assert.Equal(t, len(Hash16()), 16)
}

func TestHash32(t *testing.T) {
	assert.Equal(t, len(Hash32()), 32)
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
