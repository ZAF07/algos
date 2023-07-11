package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
	// nums := []int{4, 5, 6, 0, 1, 2, 3} // len = 6 | max = 15
	// nums := []int{4, 5, 6, 10, 120, 2, 3} // len = 6 | max = 141
	nums := []int{4, 5, 3, 4, 3} // len = 6 | max = 142
	// nums := []int{2, 4} // len = 6 | max = 142
	fmt.Println(minimumSizeSubarray(nums, 7)) // == 2
}

func minimumSizeSubarray(nums []int, target int) int {
	min := 0
	window := 1

	for window < len(nums)-1 {
		currentMin := getSum(nums[0 : 0+window]...)
		fmt.Println("currentmin : ", currentMin)
		if currentMin == target {
			if min == 0 {
				min = window
			}
		}

		for i := 1; i <= len(nums)-window; i++ {
			tempMin := getSum(nums[i : i+window]...)
			fmt.Println("tempmin: ", tempMin)
			if tempMin == target && window < min {
				min = window
			}
		}

		window++
	}

	// Check the sum of the entire array at the last step because we are searching for the MINIMAL LENGTH of subarray
	if getSum(nums...) == target {
		return len(nums)
	}

	return min
}

func getSum(i ...int) int {
	var result int
	for _, v := range i {
		result += v
	}
	return result
}
