package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
	// nums := []int{4, 5, 6, 0, 1, 2, 3} // len = 6 | max = 15
	// nums := []int{4, 5, 6, 10, 120, 2, 3} // len = 6 | max = 141
	// nums := []int{4, 50, 6, 10, 120, 2, 10} // len = 6 | max = 142
	nums := []int{2} // len = 6 | max = 142
	fmt.Println(maxAvgSum(nums, 1))
}

func maxAvgSum(nums []int, k int) int {
	if len(nums) == 1 && k == 1 {
		return nums[0]
	}
	var maxSum int
	// Get the first k elements in the nums array and calculate the sum
	maxSum = getSum(nums[0:k]...)

	for i := 1; i < len(nums)-k+1; i++ {
		// tempSum := (maxSum - nums[i-1]) + nums[i+k-1]
		tempSum := getSum(nums[i : i+k]...)
		if tempSum > maxSum {
			maxSum = tempSum
		}
	}
	return maxSum
}

func getSum(i ...int) int {
	var result int
	for _, v := range i {
		result += v
	}
	return result
}
