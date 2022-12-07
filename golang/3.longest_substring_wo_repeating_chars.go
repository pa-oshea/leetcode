package main

import (
	"math"
)

func lengthOfLongestSubstring(s string) (r int) {
	mymap := map[string]int{}
	windowStart := 0
	for index, el := range s {
		mymap[string(el)]++

		for mymap[string(el)] > 1 {
			mymap[string(s[windowStart])] -= 1
			windowStart++
		}

		r = int(math.Max(float64(r), float64(index)-float64(windowStart)+1))

	}
	return
}

// TODO: understand this better
func lengthOfLongestSubstringOptimised(s string) (r int) {
	n := len(s)
	charMap := map[string]int{}
	for i, j := 0, 0; j < n; j++ {
		if _, ok := charMap[string(s[j])]; ok {
			i = int(math.Max(float64(charMap[string(s[j])]), float64(i)))
		}
		r = int(math.Max(float64(r), float64(j-i+1)))
		charMap[string(s[j])] = j + 1
	}
	return
}
