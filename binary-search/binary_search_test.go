package binarysearch

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	For Binary Search to work, the slice has to be sorted.

	Given an array which is sorted, find the index at which the target is located at.

	Example 1:
	Input: [1,2,3,4,5], target = 5
	Output: 4
	Solution: The target 5 is located at index 5 of the array

	Example 2:
	Input: [1,2,3,4,5], target = 2
	Output: 1
	Solution: The target 2 is located at index 1 of the array
*/

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{10, 21, 33, 43, 51}, 21, 1},
		{[]int{1, 4, 9, 21, 121}, 23, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := BinarySearch(tt.nums, tt.target)
			t.Logf("RETURNED: %v", got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BinarySearch(arr []int, target int) (result int) {
	left := 0
	right := len(arr) - 1

	// As long as the left pointer comes before the right pointer
	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			result = mid
			return
		}

		// Curr is lesser than target
		if arr[mid] < target {
			left = mid + 1
		}
		// Curr is larger than target
		if arr[mid] > target {
			right = mid - 1
		}
	}
	result = -1
	return
}
