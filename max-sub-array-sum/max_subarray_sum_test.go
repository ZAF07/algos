package maxsubarraysum

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// import "fmt"

// // 🚨 NOT COMPLETE. GO SEARCH FOR THE ANSWERS

/* 53 Maximun Subarray
Given an integer array nums, find the
subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{2, -3, 4, -2, 2, 5}, 9},
		{[]int{4, -3, 4, -2, 2, -6, 1}, 5},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{10, 9, 1, -21, 20, -2}, 20},
		{[]int{-3, -2, 0, -1}, 0},
		{[]int{-2, -1}, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v | Want: %v", tt.input, tt.want)
		t.Run(testName, func(t *testing.T) {

			got := maxSubArraySum(tt.input)
			assert.Equal(t, tt.want, got)

			got1 := maxSubArraySum2(tt.input)
			assert.Equal(t, tt.want, got1)
		})
	}
}

func maxSubArraySum(nums []int) int {
	// If there is only one integer in the input, we simply return it
	if len(nums) == 1 {
		return nums[0]
	}

	// Variable to store the current subarray's sum
	cur := 0

	// Variable to store the max subarray's sum
	max := nums[0]

	for _, n := range nums {
		// If cur < 0 means that we have reached the end of a possible subarray. Meaning that we started somewhere in the array up until this point where current has become a negative number
		if cur < 0 {
			cur = 0
		}

		// add to current sum and set the maximum value to the result
		cur += n
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}

func maxSubArraySum2(nums []int) int {
	maxSum := math.MinInt
	curSum := 0

	for _, num := range nums {
		curSum = int(math.Max(float64(num), float64(curSum+num)))
		maxSum = int(math.Max(float64(curSum), float64(maxSum)))
	}
	return maxSum
}
