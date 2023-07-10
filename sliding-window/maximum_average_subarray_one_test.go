package slidingwindow

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000
*/

func TestMaximumAverageSubarrayOne(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{11, 2, 9, 4, 51}, 4, 66},
		{[]int{11, 2, 9, 4, 51, 41, 3, 5}, 2, 92},
		{[]int{3, 2, 1, 9}, 2, 10},
		{[]int{2, 4}, 1, 4},
		{[]int{2}, 1, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given : %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := maxAvgSum(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func maxAvgSum(nums []int, k int) int {
	// Return early if the length of the array is the same as k
	if len(nums) == 1 && k == 1 {
		return nums[0]
	}

	var maxSum int
	// Get the first k elements in the nums array and calculate the sum
	maxSum = getSum(nums[0:k]...)

	// Loop Kx lesser than the the len of the array, bnecause we are including Kx elements after nums[i] so we don't go out of range
	for i := 1; i <= len(nums)-k; i++ {
		// Get the sum of the current window
		tempSum := getSum(nums[i : i+k]...)
		// Compare the sum of the current window to the current max sum. Replace max sum if current sum is larger (or smaller, depending on the goal)
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
