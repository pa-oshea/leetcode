package main

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	assertEqual(t, isPowerOfTwo(16), true, "16")
	assertEqual(t, isPowerOfTwo(64), true, "64")
	assertEqual(t, isPowerOfTwo(-64), false, "-64")
	assertEqual(t, isPowerOfTwo(-16), false, "-16")
}
