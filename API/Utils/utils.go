package utils

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	math_rand "math/rand"
	"os"
	"strconv"
)

var urlSize int

func init() {
	var byt [8]byte
	_, err := crypto_rand.Read(byt[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(byt[:])))

	size, ok := os.LookupEnv("SHORT_URL_SIZE")
	if !ok {
		urlSize = 8
		return
	}
	urlSize, err = strconv.Atoi(size)
	if err != nil {
		fmt.Println("error while converting string to int size", size)
	}
}

func randomBytes() []byte {
	buf := make([]byte, 8)
	_, err := crypto_rand.Read(buf)
	if err != nil {
		panic("cannot generate random bytes")
	}
	return buf
}

func GetRandomShortUrl() string {
	data := randomBytes()
	encoded := base58.Encode(data)
	return encoded[:urlSize]
}
