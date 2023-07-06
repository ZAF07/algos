package rotatearray

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

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
