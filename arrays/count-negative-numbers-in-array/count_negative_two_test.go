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

func CountNegativeIntegersTwo(grid [][]int) int {
	res := 0
	for i := 0; i <= len(grid)-1; i++ {
		curArray := grid[i]
		curNum := curArray[len(curArray)-1]

		for curNum < 0 {
			fmt.Println("a: ", curArray)
			if len(curArray) > 1 {
				curArray = curArray[:len(curArray)-1]
				curNum = curArray[len(curArray)-1]
			} else {
				res++
				break
			}
			res++
		}
	}
	return res
}
