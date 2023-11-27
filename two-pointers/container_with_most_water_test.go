package twopointers

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
11. Container With Most Water
Medium
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

Constraints:
n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		given []int
		want  int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := maxArea(tt.given)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

func maxArea(height []int) int {
	var maxAmt int
	l, r := 0, len(height)-1

	for l < r {
		// get the max width of the current container
		width := (r - l)

		// find the minimum height of the current container
		curHeight := int(math.Min(float64(height[l]), float64(height[r])))

		// calc the max amount of water the current container can hold (multiply the width by the min height of current container)
		curAmt := width * curHeight

		maxAmt = int(math.Max(float64(maxAmt), float64(curAmt)))

		// Now find other possible answers by moving the pointers based on conditions
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxAmt
}
