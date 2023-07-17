package slidingwindow

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

Example 1:
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000
*/

func TestMaximumAverageSubarrayOne(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{11, 2, 9, 4, 51}, 4, 16.5},
		{[]int{11, 2, 9, 4, 51, 41, 3, 5}, 2, 46},
		{[]int{3, 2, 1, 9}, 2, 5.0},
		{[]int{2, 4}, 1, 4.0},
		{[]int{2}, 1, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given : %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := maxAvgSum(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func maxAvgSum(nums []int, k int) float64 {
	// Return early if nums has only one element
	if len(nums) < 2 && k == 1 {
		return float64(nums[0])
	}
	// This is something i have to check. I don't quite understand this. Everyting does not work if i don't init sum this way
	sum := float64(-1 << 63)
	cSum := float64(0)
	left, right := 0, 0
	// Loop len(nums) times to capture all possible window size
	for right <= len(nums)-1 {
		// We keep adding to the result the value of the right (Window grows to the right)
		cSum += float64(nums[right])

		// Here we check if the size of the window is == k && that the right pointer is still in range (within the length of the array)
		// This means that we have reached a valid window size and have to calculate the sum of the current window
		cLen := (right - left) + 1
		if cLen == k && right <= len(nums)-1 {
			sum = math.Max(cSum, sum)
			// We substract the left element from the sum here (we are sliding the window of size k to the right side. So we have to remove the left most element first)
			cSum -= float64(nums[left])
			left++
		}
		// At the end of each pass, we always add one to the right pointer to keep moving the sliding window. We only adjust the left pointer when the window has reached it max size
		right++
	}
	return (sum / float64(k))
}
