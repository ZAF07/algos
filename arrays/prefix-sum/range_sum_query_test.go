package prefixsum

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type TestCase struct {
	Left, Right int
	Want        int
}

func TestRangeSumQuery(t *testing.T) {
	tests := []struct {
		Arr       []int
		TestCases []TestCase
	}{
		{[]int{1, 2, 3, 4}, []TestCase{
			{Left: 0, Right: 2, Want: 6},
			{Left: 0, Right: 3, Want: 10},
			{Left: 1, Right: 2, Want: 5},
			{Left: 1, Right: 3, Want: 9},
			{Left: 2, Right: 3, Want: 7},
			{Left: 3, Right: 3, Want: 4},
		}},
		{[]int{4, 11, 3, 5}, []TestCase{
			{Left: 0, Right: 2, Want: 18},
			{Left: 0, Right: 3, Want: 23},
			{Left: 1, Right: 2, Want: 14},
			{Left: 1, Right: 3, Want: 19},
			{Left: 2, Right: 3, Want: 8},
			{Left: 3, Right: 3, Want: 5},
		}},
	}

	for _, tt := range tests {
		nums := NewNumArr(tt.Arr)
		for _, tc := range tt.TestCases {
			testName := fmt.Sprintf("Given arr: %+v || Left: %d & Right: %d || Want: %d", tt.Arr, tc.Left, tc.Right, tc.Want)
			t.Run(testName, func(t *testing.T) {
				got := nums.SumRange(tc.Left, tc.Right)
				want := tc.Want
				assert.Equal(t, want, got)
			})

		}
	}
}

type NumArray struct {
	Nums []int
}

func NewNumArr(nums []int) NumArray {
	// Calc the prefix sum up until the current idx
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{nums}
}

func (n *NumArray) SumRange(l, r int) int {
	if l == 0 {
		return n.Nums[r]
	}
	return n.Nums[r] - n.Nums[l-1]
}
