package slidingwindow

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

func TestMinSizeSubarraySum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 2, 4, 3, 6}, 7, 2},
		{[]int{4, 2, 4, 3, 6}, 6, 1},
		{[]int{1, 1, 1, 1, 1, 1}, 6, 6},
		{[]int{1}, 6, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := minimumSizeSubarray(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func minimumSizeSubarray(nums []int, target int) int {
	left := 0
	right := 0
	sum := 0
	min := math.Inf(1)

	for right < len(nums) {
		sum += nums[right]

		for sum >= target {
			size := right - left + 1
			min = math.Min(float64(size), min)

			sum -= nums[left]
			left++
		}
		right++
	}

	if min == math.Inf(1) {
		return 0
	}
	return int(min)
}
