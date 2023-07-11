package twosum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
  Find the two numbers that when added together, makes the target value
  Example: Target = 9, Args = [2,4,6,7,8]
  Answer: [0, 1] Indexes of 2 & 7 (2 + 7 = 9)

  NOTES: Hash performance is always faster albeit slightly
*/

func TestTwoSum(t *testing.T) { // Run 'go test -test.v' to see logs in stdout
	var testCases = []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3, 4, 5}, 7},
		{[]int{12, 2, 30, 3, 5}, 7},
		{[]int{9, 2, 2, 4, 52}, 61},
		{[]int{10, 1, 5, 2, 9}, 12},
		{[]int{8, 2, 4, 2, 11}, 19},
	}

	for _, tt := range testCases {
		testName := fmt.Sprintf("Given %v, Expected: %v", tt.nums, tt.target)
		t.Run(testName, func(t *testing.T) {
			got := 0
			res := TwoSum(tt.nums, tt.target)
			t.Log("Returned slice from TwoSum(): ", res)
			for _, v := range res {
				got += tt.nums[v]
			}
			assert.Equal(t, got, tt.target)

		})
	}

}

func TwoSum(nums []int, target int) []int {
	// Store the missin pair needed to make up target
	pair := make(map[int]int)

	for k, v := range nums {
		// Check if the missing pair for the current value is in the hash map
		if val, ok := pair[v]; ok {
			// Return the two indexes of the pair that makes up the target
			return []int{k, val}
		}
		// Store in the hash map the missing pair needed to make the target
		// The key would be the value of the missing pair (target - current value (num[i]) )
		// The value would be the index of the array where we can find the current pair
		pair[target-v] = k
	}
	return []int{}
}
