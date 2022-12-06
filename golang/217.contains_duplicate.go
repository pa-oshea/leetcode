package main

func containsDuplicate(nums []int) (r bool) {
	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++
		if counter[v] > 1 {
			return true
		}
	}
	return
}
