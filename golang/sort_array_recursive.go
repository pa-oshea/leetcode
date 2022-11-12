package main

func sort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	temp := arr[len(arr)-1]
	arr = remove(arr, len(arr)-1)

	sort(arr)

	arr = insert(arr, temp)

	return arr
}

func insert(arr []int, temp int) []int {
	if len(arr) == 0 || temp >= arr[len(arr)-1] {
		arr = append(arr, temp)
		return arr
	}

	v := arr[len(arr)-1]
	arr = remove(arr, len(arr)-1)

	arr = insert(arr, temp)

	arr = append(arr, v)

	return arr
}