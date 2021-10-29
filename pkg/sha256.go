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