package main

import (
	"fmt"
)

func search(array []int, target int) int {
	min := 0
	max := len(array) - 1
	mid := (min + max) / 2

	for min <= max {

		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			min = mid + 1
			mid = (min + max) / 2
		} else {
			max = mid - 1
			mid = (min + max) / 2
		}
	}
	return 0
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 6, 7}, 4))
}
