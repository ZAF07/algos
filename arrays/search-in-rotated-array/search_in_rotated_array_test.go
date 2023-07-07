package searchinrotatedarray

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestSearchInRotatedArray(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 1, 5},
		{[]int{0, 1, 2, 4, 5, 6, 7}, 0, 0},
		{[]int{5, 4, 3, 0, 1, 2}, 4, 1},
		{[]int{5, 4, 3, 0, 1, 2}, 10, -1},
		{[]int{7, 6, 5, 0, 1, 2}, 3, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := SearchInRotatedArray(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Time: O(log n) || Space: O(1)
func SearchInRotatedArray(nums []int, target int) int {
	// First we get the index holding the smallest element in the array using the Binary Search method
	pivot := searchPivot(nums)

	// In the first pass, we pass in the right half (if array was rotated) of the array first
	res := binarySearch(nums, pivot, len(nums)-1, target)
	if res != -1 { // 💡💡 COULD WE MAKE IT SO THAT WE KNOW HERE IF THE ANSWER IS RIGHT OR WRONG SO WE DONT HAVE TO RUN THE BINARYSEARCH() METHOD ON THE OTHER HALF ??
		return res
	} else {
		// We pass in the left half in the next pass
		return binarySearch(nums, 0, pivot, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func searchPivot(nums []int) (mid int) {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid = (left + right) / 2

		if nums[mid] > nums[len(nums)-1] {
			left = mid + 1
		}

		if nums[mid] < nums[len(nums)-1] {
			right = mid - 1
		}
	}

	return mid
}
