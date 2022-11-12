package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	highestTotal := nums[0]
	for j := 0; j < len(nums); j++ {
		sumResult := nums[j]
		for i := j + 1; i < len(nums); i++ {
			sumResult += nums[i]
			// DRY - come back for better solution
			if highestTotal < sumResult {
				highestTotal = sumResult
			}
		}
		if highestTotal < sumResult {
			highestTotal = sumResult
		}
	}
	return highestTotal
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-1}))
	fmt.Println(maxSubArray([]int{-12, 1}))
	fmt.Println(maxSubArray([]int{-2, 1}))
}
