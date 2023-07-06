package main

import (
	"flag"
	"fmt"
)

/*
	For Binary Sort to work, the slice has to be sorted.
*/
func main() {
	var target int
	flag.IntVar(&target, "target", 0, "Enter the int you would want to search for in the slice.")
	flag.Parse()
	fmt.Println("TARGET: ", target)
	res := binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target)
	fmt.Printf("RESULT: index position %d\n", res)
}

func binarySearch(arr []int, target int) (result int) {
	left := 0
	right := len(arr) - 1

	// As long as the left pointer comes before the right pointer
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			result = mid
			return
		}

		// Curr is lesser than target
		if arr[mid] < target {
			left = mid + 1
		}
		// Curr is larger than target
		if arr[mid] > target {
			right = mid - 1
		}
	}
	result = -1
	return
}
