package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expect := "Hello, world!"
	assert.Equal(t, expect, Hello())
}
