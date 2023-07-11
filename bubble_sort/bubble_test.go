package bubble

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	Given an unsorted array, sort the given array in ascending order

	Example 1:
	Input: [5,2,4,1,3]
	Output: [1,2,3,4,5]
	Solution: Sort the given array in ascending order
*/

func TestBubbleSort(t *testing.T) { // Run 'go test <path> -test.v' to see logs in stdout
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 9, 6, 1, 8}, []int{1, 2, 5, 6, 8, 9}},
		{[]int{5, 2, 3, 6, 1, 4}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{15, 12, 13, 16, 11, 14}, []int{11, 12, 13, 14, 15, 16}},
		{[]int{15, 2, 3, 26, 1, 45}, []int{1, 2, 3, 15, 26, 45}},
		{[]int{50, 23, 31, 62, 12, 4}, []int{4, 12, 23, 31, 50, 62}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given : %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := Bubble(tt.nums)
			// t.Logf("RETURNED: %v", got)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Bubble(arr []int) []int {
	// outer pass goes through the entire length of the array
	for i := 0; i < len(arr); i++ {
		// inner pass goes through the entire length of the array as well -i -1
		for j := 0; j < len(arr)-1-i; j++ { // We always decrease the inner loop by i so that we dont visit the last element in each pass. Because we already know that after each pass, the largest element encountered will be pushed to the back of the slice
			// check if current array element is larger then the element right next to it
			if arr[j] > arr[j+1] {
				// if so, swap the current element with the element right next to it
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

// Slight convenience , slightly confusing for others to read ...
func bubbleTwo(nums []int) []int {
	// compare current and next value, if next is > current, swap positions
	for i := 0; i < len(nums); i++ { // First pass goes through the entire array.
		for j := 0; j > len(nums)-1-i; j++ { // inner loop keeps decreasing because we know that at each pass, the last element will be the largest element
			if nums[j] > nums[j+1] { // if the current num is > next num, we swap positions
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
