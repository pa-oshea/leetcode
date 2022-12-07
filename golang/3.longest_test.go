package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
}

func TestLengthOfLongestSubstringOptimised(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstringOptimised("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("pwwkew"))
	assert.Equal(t, 3, lengthOfLongestSubstringOptimised("dvdf"))
}
