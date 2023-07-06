package main

import "fmt"

/*
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

Example 1:

Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
Output: 8
Explanation: There are 8 negatives number in the matrix.
*/

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	res := countNegativeNumbers(grid)
	fmt.Println(res)
}

func countNegativeNumbers(grid [][]int) (count int) {
	lengthOfGrid := len(grid)
	for i := 0; i < lengthOfGrid; i++ {
		// I start from the last element of the slice because we know that the slice is sorted decreasingly
		lengthOfRow := len(grid[i]) - 1
		for j := lengthOfRow; j >= 0; j-- {
			cur := grid[i][j]
			// If cur is a negative number, we increment count by 1
			if cur < 0 {
				count++
				continue
			}
			// Break the inner loop if we hit a non negative number so that we dont habe to visit all elements
			break
		}
	}
	return
}
