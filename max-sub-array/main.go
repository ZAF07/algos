package main

import "fmt"

// 🚨 NOT COMPLETE. GO SEARCH FOR THE ANSWERS

func main() {
	nums := []int{-1, 1, 2, 3, 4, -1, 5, 5, 5, 5, 5, 5, 5}
	res := maxSubArray(nums)
	fmt.Println("Result: ", res)
}

func maxSubArray(nums []int) int {
	curSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(curSum+nums[i], nums[i])
		maxSum = max(curSum, maxSum)
		fmt.Println("maxsum: ", maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
