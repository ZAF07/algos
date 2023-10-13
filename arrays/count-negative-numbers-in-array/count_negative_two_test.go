package countnegativenumbersinarray

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestCountNegativeTwo(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 8},
		{[][]int{{4, 3, -2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}, 9},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v | Want: %v", tt.grid, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := CountNegativeIntegersTwo(tt.grid)
			assert.Equal(t, got, tt.want)
		})
	}
}

/*
The runtime for this algo is O(n * m) best case only when n > m.
When m > n, this algo is worst than O(n^2)
*/
func CountNegativeIntegersTwo(grid [][]int) int {
	res := 0
	for i := 0; i <= len(grid)-1; i++ {
		curArray := grid[i]
		curNum := curArray[len(curArray)-1]

		// As long as the current integer is a negative number, we continuw with the loop
		for curNum < 0 {
			// Imcrement the result at each loop. (This only happens when current number is a negative number)
			res++
			if len(curArray) > 1 {
				// Pop the last element from the current array
				curArray = curArray[:len(curArray)-1]
				// Update the current number
				curNum = curArray[len(curArray)-1]
			} else {
				break
			}
		}
	}
	return res
}
