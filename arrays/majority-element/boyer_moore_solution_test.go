package majorityelement

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestBoyerMooreMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1}, 1},
		{[]int{3, 1, 3, 2, 1, 1}, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Boyer-Moore Voting Algorithm Implementation...
func majorityElement(nums []int) int {
	count := 1
	candidate := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
		}

		if count == 0 {
			candidate = nums[i]
			count = 1
		}
	}
	return candidate
}
