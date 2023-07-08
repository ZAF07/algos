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
		// {[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 11, -1},
		// {[]int{0, 1, 2, 3, 4, 5, 6, 7}, 1, 1},
		// {[]int{5, 0, 1, 2, 3, 4}, 0, 1},
		// {[]int{3, 4, 5, 0, 1, 2}, 3, 0},
		// {[]int{6, 7, 0, 1, 2, 3}, 3, 5},
		// {[]int{2, 3, 4, 5, 6, 7, 8, 1}, 3, 1},
		// {[]int{3, 1, 2}, 3, 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			// got := SearchInRotatedArray(tt.nums, tt.target)
			got := SearchInRotatedArray2(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Time: O(log n) || Space: O(1)
func SearchInRotatedArray(nums []int, target int) int {
	// First we get the index holding the smallest element in the array using the Binary Search method
	pivot := findPivot(nums)

	if target > nums[len(nums)-1] {
		return binarySearch(nums, 0, pivot, target)
	}
	return binarySearch(nums, pivot, len(nums)-1, target)
}

//  3,1,2
func binarySearch(nums []int, left, right, target int) int {
	fmt.Println("GOTTENT IN BINARY : ", nums, left, right)
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

//  The pivot is the smallst element in the array.
func findPivot(nums []int) int {
	var mid int
	left := 0
	right := len(nums) - 1
	lastEl := nums[len(nums)-1]

	for left <= right {
		mid = (left + right) / 2

		if left == right {
			fmt.Println("PIVOT: ", mid)
			return mid
		}

		if nums[mid] > lastEl {
			left = mid + 1
			continue
		}

		if nums[mid] < lastEl {
			right = mid - 1
			continue
		}

		if nums[mid] == lastEl {
			if nums[mid-1] < nums[mid] {
				left = mid - 1
				continue
			}
		}

	}
	fmt.Println("PIVOT: ", mid)
	return mid
}

//  3,1,2
func SearchInRotatedArray2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		fmt.Println("mid, left, right: ", mid, left, right)

		//  Since we are already at the altered left. right and mid position, we can simply return the values if match
		if nums[mid] == target {
			fmt.Println("found in mid")
			return mid
		}
		if nums[right] == target {
			fmt.Println("found in right")
			return right
		}
		if nums[left] == target {
			fmt.Println("found in left")
			return left
		}

		//  Go left
		if nums[left] < nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
			//  Go right
		} else {
			if target < nums[mid] || target > nums[left] {
				left = mid - 1
			} else {
				right = mid + 1
			}
		}
	}
	return -1
}
