package twosumtwo

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
167. Two Sum II - Input Array Is Sorted
Medium

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.

Example 1:
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Constraints:
2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.
*/

func TestTwoSumTwo(t *testing.T) {
	tests := []struct {
		given  []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 4}},
		{[]int{10, 20, 34, 41, 51}, 75, []int{3, 4}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Target: %v || Want: %v", tt.given, tt.target, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := twoSumTwo(tt.given, tt.target)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

/*
Because nums is sorted, we can simply loop the array and increase/decrease the pointers based on whether sum is larger/smaller than target
*/
func twoSumTwo(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		cur := nums[l] + nums[r]

		if cur == target {
			// because we are guaranteed one solution, we can break and return at the end of the loop. no need for default values
			break
		}

		if cur < target {
			l++
		}

		if cur > target {
			r--
		}
	}

	return []int{l + 1, r + 1}
}
