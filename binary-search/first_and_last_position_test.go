package binarysearch

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/* 34. Find First and Last Position of Element in Sorted Array
Medium

Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.
If target is not found in the array, return [-1, -1].
You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109
*/

func TestFirstAndLast(t *testing.T) {
	tests := []struct {
		target int
		given  []int
		want   []int
	}{
		{2, []int{1, 1, 2, 2, 2, 3, 4, 5}, []int{2, 4}},
		{4, []int{4, 2, 2, 3, 3, 4, 4}, []int{5, 6}},
		{2, []int{2}, []int{0, 0}},
		{0, []int{}, []int{-1, -1}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := fnl2(tt.given, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

// O(n) version
func FnL(nums []int, target int) []int {
	res := make([]int, 2)

	// Find the start Idx
	l, r := 0, len(nums)-1
	start := -1

	for l <= r {
		mid := (l + r) / 2
		// 1,1,2,2,3,3
		if nums[mid] == target {
			// Find the first instance of the target
			for mid-1 >= 0 && nums[mid-1] == target {
				mid--
			}
			start = mid
			break
		}

		if nums[mid] > target {
			r = mid - 1
		}

		if nums[mid] < target {
			l = mid + 1
		}

		start = r
	}

	if start == -1 || nums[start] != target {
		return []int{-1, -1}
	}

	// Find the start and end of the target in the array
	res[0] = start
	for start < len(nums) && nums[start] == target {
		res[1] = start
		start++
	}
	return res
}

// O(logn) version
func fnl2(nums []int, target int) []int {
	start := findTarget(nums, target, true)
	if start == -1 {
		return []int{-1, -1}
	}
	end := findTarget(nums, target, false)
	return []int{start, end}
}

/*
Modified binary search.

We basically perform a regular binary search.

Difference with this version is that when we find the target value, we keep the pointers moving (direction depends on whether we're looking for the first or last instance of the target value)

Eventually, left pointer will go past the right pointer, which is when the binary search terminates.
*/
func findTarget(nums []int, target int, findMin bool) int {
	l, r := 0, len(nums)-1
	idx := -1
	for l <= r {
		mid := (l + r) / 2

		//  if we are looking for the first occurance, we shrink the window to the left, else shrink to the right
		if nums[mid] == target {
			if findMin {
				r = mid - 1
			} else {
				l = mid + 1
			}
			idx = mid
		}

		if nums[mid] > target {
			r = mid - 1
		}

		if nums[mid] < target {
			l = mid + 1
		}

	}
	return idx
}
