package removeduplicates

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	Task is to remove any duplicate element in the array and return the number of remaining elements in the array
	Example:
			[1,1,2] > [1,2] == 2
			[2,3,1,4,2] > [2,3,1,4] == 4

	💡This runs at 0(n) time complexity
*/

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 1, 1}, 2},
		{[]int{0, 2, 1, 2}, 3},
		{[]int{0, 2, 1, 2, 3, 4, 1, 2, 4}, 5},
		{[]int{6, 3, 1, 2}, 4},
		{[]int{2, 2, 1, 12}, 3},
		{[]int{1, 1, 1, 2}, 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Ecpected : %v", tt.want)
		t.Run(testName, func(t *testing.T) {
			// got := RemoveDuplicates(tt.nums)
			got := RemoveDuplicates2(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

/* Implementation #1
Here we have a variable that stores the index of the array starting at position 0
At every iteration, if we have not seen the value, we add it to the map and adjust the array at current position to the current element
We only increase the position if we have not seen the current value and after we have added the current value to the array

*/
func RemoveDuplicates(arr []int) int {
	seen := make(map[int]int)
	idx := 0
	for _, val := range arr {

		// If the current value is NOT seen/visited in the loop yet
		if _, ok := seen[val]; !ok {

			// add value to the map for future reference
			seen[val] = val

			// set the slice at current index to the value
			arr[idx] = val

			// Increment idx only when the value has not been seen yet(in other words, idx remains the same if the value has been seen before) so that we can replace that value with a value that has NOT been seen. With this we know that any values after arr[idx] has been seen before so we can trim the slice(arr[:idx]) up to the idx and return the length as the result
			idx++
		}
	}
	arr = arr[:idx]
	return len(arr)
}

/*  Implementation #2
This method loops the given array once.
At each iteration, the array becomes smaller if ther are any duplicates
*/
func RemoveDuplicates2(nums []int) int {
	seenMap := make(map[int]int)

	// Loop the array starting from the end
	for i := len(nums) - 1; i >= 0; i-- {
		// If we've seen the value before, adjust the array removing the current seen element
		if position, seen := seenMap[nums[i]]; seen {
			nums = append(nums[:position], nums[position+1:]...)
		}

		// Regardless of whether we've seen the current element or not, we update its latest seen position
		seenMap[nums[i]] = i

	}
	return len(nums)
}
