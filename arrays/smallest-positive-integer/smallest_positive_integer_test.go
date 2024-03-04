package smallestpositiveinteger

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	Given an array of integers ranging from negative to positive, return the smallest non-negative integer that does not appear in the array.

	Example:

	Given [-1,-2,4,2,3,5], Answer: 2
*/

func TestSmallestPoositiveInteger(t *testing.T) {
	tests := []struct {
		given []int
		want  int
	}{
		{[]int{-1, -2, -3, 1}, 2},
		{[]int{4, 3, 1, 5, -1}, 2},
		{[]int{2, 6, 4, 7, 4}, 1},
		{[]int{1, 2, 6, 4, 7, 4}, 3},
		{[]int{-1, 1, 2, -8, 0, 6, 4, 7, 4}, 3},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v || Want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := smallestPositiveInt(tt.given)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

func smallestPositiveInt(nums []int) int {
	// store positive numbers in given num
	seen := make(map[int]bool)

	// iterate given nums and add to map only if not seen && is positive number
	for _, num := range nums {
		if !seen[num] {
			if num > 0 {
				seen[num] = true
			}
		}
	}

	// start looping from 1 (smallest positive number) and return the first positive number that does not appear in map
	for i := 1; i < math.MaxInt; i++ {
		if !seen[i] {
			return i
		}
	}
	return -1
}
