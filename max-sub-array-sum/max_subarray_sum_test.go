package maxsubarraysum

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// import "fmt"

// // 🚨 NOT COMPLETE. GO SEARCH FOR THE ANSWERS

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{2, -3, 4, -2, 2, 5}, 9},
		{[]int{4, -3, 4, -2, 2, -6, 1}, 5},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v | Want: %v", tt.input, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := maxSubArraySum(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func maxSubArraySum(nums []int) int {
	cur := 0
	max := nums[0]

	if len(nums) == 1 {
		return nums[0]
	}

	for _, n := range nums {
		if cur < 0 {
			cur = 0
		}

		cur += n
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}
