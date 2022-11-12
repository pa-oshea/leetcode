package main

import (
	"fmt"
	"testing"
)

func TestPyramid(t *testing.T) {
	t.Log("hello")
	fmt.Println(Pyramid(3))
	// assertEqual(t, Pyramid(3), [][]int{[ [1], [1,1], [1,1,1] ]}, "")
}
