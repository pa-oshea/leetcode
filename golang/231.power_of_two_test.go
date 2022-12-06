package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, true, isPowerOfTwo(16), "16")
	assert.Equal(t, true, isPowerOfTwo(64), "64")
	assert.Equal(t, false, isPowerOfTwo(-64), "-64")
	assert.Equal(t, false, isPowerOfTwo(-16), "-16")
}
