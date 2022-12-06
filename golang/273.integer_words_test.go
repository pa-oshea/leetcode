package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberToWords(t *testing.T) {

	assert.Equal(t, "Ten", numberToWords(10), "")
	assert.Equal(t, "One Hundred Twenty Three", numberToWords(123), "")
	assert.Equal(t, "One Thousand", numberToWords(1000), "")
	assert.Equal(t, "Ten Thousand", numberToWords(10000), "")
	assert.Equal(t, "One Hundred Thousand", numberToWords(100000), "")
	assert.Equal(t, "One Million", numberToWords(1000000), "")
	assert.Equal(t, "Twelve Thousand Three Hundred Forty Five", numberToWords(12345), "")
	assert.Equal(t, "Fifty Thousand Eight Hundred Sixty Eight", numberToWords(50868), "")
	assert.Equal(t, "Fifty Eight Thousand Eight Hundred Sixty Eight", numberToWords(58868), "")
	assert.Equal(t, "One Billion", numberToWords(1000000000), "")

}
