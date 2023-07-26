package twosum

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// Same as TwoSum but this one epects an array of the numbers that make up the target value

func TestTwoSumArray(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{[]int{4, 3}, 10, [][]int{}},
		{[]int{4, 3, 4, 3}, 7, [][]int{{3, 4}}},
		{[]int{3, 3, 3, 3}, 6, [][]int{{3, 3}}},
		{[]int{3, 0, 2, 1}, 3, [][]int{{0, 3}, {1, 2}}},
		{[]int{0, 3, 4, 3, 2, 6}, 6, [][]int{{3, 3}, {2, 4}, {6, 0}}},
		{[]int{0, 3, 4, 3, 2, 6, 3, 3, 3}, 6, [][]int{{3, 3}, {2, 4}, {6, 0}}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := twoSumArray(tt.nums, tt.target)

			assert.Equal(t, tt.want, got)
		})
	}
}

func twoSumArray(nums []int, target int) [][]int {
	res := [][]int{}
	seenMap := make(map[int]int)

	for i, num := range nums {
		missingNum := target - num
		if val, seen := seenMap[num]; seen {
			if val == -1 {
				continue
			}
			res = append(res, []int{num, nums[val]})
			seenMap[num] = -1
			continue
		}

		if _, ok := seenMap[missingNum]; ok {
			continue
		}
		seenMap[missingNum] = i

	}
	return res
}
