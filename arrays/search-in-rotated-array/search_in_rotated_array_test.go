package searchinrotatedarray

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Example 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Example 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Example 3:
Input: nums = [1], target = 0
Output: -1

*/

func TestSearchInRotatedArray(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, -1, -1},
		{[]int{3, 1, 2}, 3, 0},
		{[]int{3, 1, 2}, 4, -1},
		{[]int{3, 1, 2}, 0, -1},
		{[]int{2, 1}, 1, 1},
		{[]int{2, 1}, 4, -1},
		{[]int{2, 1}, 0, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := SearchInRotatedArray(tt.nums, tt.target)
			// got := SearchInRotatedArray2(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Time: O(log n) || Space: O(1)
func SearchInRotatedArray(nums []int, target int) int {
	// First we find the pivot point of the rotated array
	// This is so we can perform Binary Search (Which works well with a sorted array) on either the right/left half
	pivot := findPivot(nums)

	// if pivot is -1, this means that the given array is not rotated, so we can simply pass in the entire array and perform Binary Search on it
	if pivot == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	// Remember that the given array is sorted in ascending order (elements to the right are always greater than left unless rotated) AND rotated, so if the target is < nums[0], we can start at the pivot point (the smallest element in the given array)
	if target < nums[0] {
		return binarySearch(nums, pivot, len(nums)-1, target)
		// If the target is > than nums[0], we pass in the array from nums[0] up untill the pivot -1 (pivot is the smallest possible element in the array, so the element that comes right before it must the the largest)
	} else {
		return binarySearch(nums, 0, pivot-1, target)
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

// Find the pivot index (the point where the array is rotated aka the smallest element in the array) if it exists, return -1 otherwise
func findPivot(nums []int) int {
	// Since all elements in the array are unique,
	// if the array is rotated, the first element will be greater than the last element.
	// Hence we can do an early return if the first element is less than the last element.
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	// Binary search for the pivot index
	// The element at the pivot index will be the smallest element in the array
	ref := nums[0]

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		//💡WHY DO WE BRING THE LEFT POINTER UP IF NUM[MID] >= NUMS[0]??
		// Because we know that the given array is already sorted in ascending order and we are trying to find the smallest possible element in the given array, when a rotation happens within an ordered array, the smallest element (the first element in an unsorted array) will be moved to the right k number of times. So if the middle element is larget than the first element in a rotated array, we know that the smallest possible element will definitely be somewhere to the right
		if nums[mid] >= ref {
			left = mid + 1
		}
		// 💡WHY DO WE BRING THE RIGHT POINTER DOWN IF NUMS[MID] < NUMS[0] ??
		// Because we are trying to find the smallest possible element here (the pivot or where we can split the rotated array into two to easily perform Binary Search on afterwards) know that the array is already sorted in ascending order, any number to the right of an element will only grow larger (unless the array is rotated of course!)
		// So if the current middle element (nums[(left + right) / 2]) is lesser than the first element (nums[0]), knowing that elements to the right are only going to be larger, there is the possibility that num[mid + 1] could be larger than or equal to nums[0], so we can safely narrow down the search for the smallest element from the current left pointer up until the middle element
		if nums[mid] < ref {
			right = mid
		}
	}

	return left
}
