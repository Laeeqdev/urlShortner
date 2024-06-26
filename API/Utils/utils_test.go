package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("SHORT_URL_SIZE", "3")
}

func TestGetRandomShortUrl(t *testing.T) {
	a := GetRandomShortUrl()
	assert.NotNil(t, a)
}

func BenchmarkGetRandomShortUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomShortUrl()
	}
}
