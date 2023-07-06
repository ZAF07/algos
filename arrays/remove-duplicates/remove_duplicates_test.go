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
			got := RemoveDuplicates(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func RemoveDuplicates(arr []int) int {
	fmt.Println("Given slice: ", arr)
	seen := make(map[int]int)
	idx := 0
	for _, val := range arr {
		fmt.Println("############START#############")

		fmt.Println("Current idx is at: ", idx)

		fmt.Printf("Checking if %d has been seen before\n", val)
		// If the current value is NOT seen/visited in the loop yet
		if _, ok := seen[val]; !ok {
			fmt.Printf("Value: %d has NOT been seen\n", val)

			fmt.Printf("Adding value: %d to the seen map\n", val)
			// add value to the map for future reference
			seen[val] = val

			fmt.Printf("Setting slice element at Index:%d to Value:%d\n", idx, val)
			// set the slice at current index to the value
			arr[idx] = val
			fmt.Println("This is the slice after setting: ", arr)

			fmt.Printf("Incrementing idx from %d to %d\n", idx, idx+1)
			// Increment idx only when the value has not been seen yet(in other words, idx remains the same if the value has been seen before) so that we can replace that value with a value that has NOT been seen. With this we know that any values after arr[idx] has been seen before so we can trim the slice(arr[:idx]) up to the idx and return the length as the result
			idx++
		} else {
			fmt.Printf("Value:%d HAS been seen. Skipping\n", val)
		}
		fmt.Println("############END#############")
	}
	arr = arr[:idx]
	fmt.Println("Result: ", arr)
	return len(arr)
}
