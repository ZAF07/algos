package productofarrayexceptself

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constraints:
2 <= nums.length <= 105
-30 <= nums[i] <= 30

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

func TestProductOfArray(t *testing.T) {
	tests := []struct {
		given []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v || Want: %v", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := product(tt.given)
			want := tt.want
			assert.Equal(t, got, want)
		})
	}
}

// TODO: Study the solution
func product(nums []int) []int {
	res := make([]int, len(nums))

	// set initial value to 1 because multiplying any number by 1 returns the same number
	buffer := 1

	// start from the start of the array first
	for i := 0; i < len(nums); i++ {
		// Set the current position's value to whatever the buffer is (before multiplying the current value)
		res[i] = buffer
		// after setting the value of the current position in the result, we multiply the current value by the buffer (this is so we skip the product of the current value)
		buffer *= nums[i]
	}

	buffer = 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] *= buffer
		buffer *= nums[j]
	}

	return res
}
