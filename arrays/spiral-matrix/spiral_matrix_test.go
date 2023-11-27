package spiralmatrix

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
54. Spiral Matrix
Medium
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

/*
	This solution mimics exactly what the problem tells us to do.
	It literally goes around the matrix in a spiral
	This is done inside a 'for loop' that runs as long as the length of the result array is < len(matrix)*len(matrix[0]) (which is basically the number of elements the the entire matrix)

	The spiral goes in cycles

	Each spiral cycle ends when it completes a square

	We keep count of the spiral cycle so we know where to start/stop i & j indexes
*/

func TestSpiralMatrix(t *testing.T) {
	tests := []struct {
		given [][]int
		want  []int
	}{
		{[][]int{{1, 2, 3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := spiralMatrix(tt.given)
			want := tt.want

			assert.Equal(t, got, want)
		})
	}
}

func spiralMatrix(matrix [][]int) []int {
	// If there is only one level in the matrix array return the array
	if len(matrix) == 1 {
		return matrix[0]
	}

	var res []int

	// Basically the total number of elements inside the matrix
	matrixSize := len(matrix) * len(matrix[0])

	// This cycle determines where i & j will start or stop in the loop
	cycle := 1

	for len(res) < matrixSize {

		// Start points for i & j index
		i, j := cycle-1, cycle-1

		/*
			💡 Inside this loop, before we loop for each dorection, its condition should also include that the length of thre result is < matrix size
			so that we could simply skip the step and quickly exit the loop once we hit the last element of the matrix
		*/

		// go right
		for len(res) < matrixSize && i == cycle-1 && j <= len(matrix[i])-cycle {
			res = append(res, matrix[i][j])
			j++
		}

		// Adjust i & j pointers beacuse after the previous loop, j index would become == len(matrix[i]) which will give an index out of bounds error
		j--
		// Adjust i index to start at the next level so we don't have duplicate element in the result array
		i++
		// go down
		for len(res) < matrixSize && j == len(matrix[0])-cycle && i <= len(matrix)-cycle {
			res = append(res, matrix[i][j])
			i++
		}

		i--
		j--
		// go left
		for len(res) < matrixSize && i == len(matrix)-cycle && j >= cycle-1 {
			res = append(res, matrix[i][j])
			j--
		}

		j++
		i--
		// go up
		for len(res) < matrixSize && j == cycle-1 && i >= cycle {
			res = append(res, matrix[i][j])
			i--
		}
		// Increment cycle so in the next iteration, i & j indexes would start at the correct positions
		cycle++
	}
	return res
}
