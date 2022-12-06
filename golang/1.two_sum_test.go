package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9), "first test")
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6), "second test")
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6), "third test")

	assert.Equal(t, []int{0, 1}, twoSumOnePass([]int{2, 7, 11, 15}, 9), "first test")
	assert.Equal(t, []int{1, 2}, twoSumOnePass([]int{3, 2, 4}, 6), "second test")
	assert.Equal(t, []int{0, 1}, twoSumOnePass([]int{3, 3}, 6), "third test")
}
