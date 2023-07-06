package rotatearray

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	Given an array of integers, you are to rotate the order of the integers k number of times to the right

	Example 1:
	Input: [1,2,3], k = 2
	Output: [2,3,1]
	Solution: 1 would shift place to the right 2 times, 2 would shift right 2 times and so will 3...
*/

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3, 5, 4, 3}, 3, []int{5, 4, 3, 1, 2, 3}},
		{[]int{1, 2, 3}, 6, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Expected: %v", tt.want)
		t.Run(testName, func(t *testing.T) {
			got := RotateArr(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func RotateArr(nums []int, k int) []int {
	x := len(nums) - k%len(nums)

	nums = append(nums[x:], nums[:x]...)
	return nums
}
