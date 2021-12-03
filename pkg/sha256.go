package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"io"
	"os"
)

func NewSHA256Client() hash.Hash {
	return sha256.New()
}

func NewBase64Client() io.WriteCloser {
	return base64.NewEncoder(base64.StdEncoding, os.Stdout)
}

func HashWithSHA256(data string) string {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	return string(hasher.Sum(nil))
}

func EncodeWithBase64(data string) string {
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer encoder.Close()

	return base64.StdEncoding.EncodeToString([]byte(data))
}
