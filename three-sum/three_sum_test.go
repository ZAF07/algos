package threesum

import (
	"fmt"
	"sort"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Medium
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:
3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func TestThreeSum(t *testing.T) {
	tests := []struct {
		given []int
		want  [][]int
	}{
		{[]int{0, 1, 1}, [][]int(nil)},
		{[]int{-2, 1, 1, 0, 1, -1}, [][]int{{-2, 1, 1}, {-1, 0, 1}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := threeSum(tt.given)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

func threeSum(nums []int) [][]int {
	// sort the array so that it is easier to check for duplicates
	sort.Ints(nums)
	var res [][]int

	// loop up until len(nums) - 2 because the left and right pointers are checking the values after i
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				l++
				r--

				// as long as the next left == the previous left value, we keep increasing l position
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				// as long as the next right == the previous right value, we keep decreasing r position
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}

			// because nums is sorted, we increase left pointer. same for right
			if sum < 0 {
				l++
			}
			if sum > 0 {
				r--
			}
		}
	}
	return res
}
