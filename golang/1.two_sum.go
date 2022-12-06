package main

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Loop through an int arrray and return the indexes of elements that sum to the target
func twoSumOnePass(nums []int, target int) []int {
	twoSumMap := map[int]int{}
	// while we are iterating through nums
	// we add to a hashmap. The key being the compliment (target - nums[i] (number we are currently looking at))
	// future loops we can look up the hashmap to see if the key (compliment) already exists
	// if it does, we return the index stored with that key and the current index we are on
	for i := 0; i < len(nums); i++ {
		// first round compliment = 9 - 2 = 7
		// second round compliment = 9 - 7 = 2
		compliment := target - nums[i]
		// first round check if the hashmap has the key of 7
		// second round check if the hashmap has the key of 2
		// which it will so return [map[2] = 0, 1]
		if _, ok := twoSumMap[compliment]; ok {
			return []int{twoSumMap[compliment], i}
		}
		// add map[2: 0]
		twoSumMap[nums[i]] = i
	}
	return nil
}
