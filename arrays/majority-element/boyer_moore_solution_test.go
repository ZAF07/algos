package majorityelement

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/* 169 Majority Element
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

Example 1:
Input: nums = [3,2,3]
Output: 3

Example 2:
Input: nums = [2,2,1,1,1,2,2]
Output: 2

Constraints:
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
*/

func TestBoyerMooreMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1}, 1},
		{[]int{3, 1, 3, 2, 1, 1, 0, 0}, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Explanation:
/*
	This works only if the Majority element is of size (n/2) or larger.
*/

// Boyer-Moore Voting Algorithm Implementation...
func majorityElement(nums []int) int {
	count := 1
	candidate := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			// For each time we see the same number, we increment count
			count++
		} else {
			// If the next number is not the same we decrement count
			count--
		}

		// Only if count == 0 do we change the candidate and reiitialise count to 1
		if count == 0 {
			candidate = nums[i]
			count = 1
		}
	}
	return candidate
}
