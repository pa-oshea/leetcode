package main

func Pyramid(n int) (result [][]int) {
	r := []int{}

	if n == 0 {
		return [][]int{}
	}

	for i := 1; i <= n; i++ {
		r = append(r, 1)
		result = append(result, r)
	}

	return
}
