package majorityelement

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	#169 Majority Element
	Given an array (nums) of size n, return the majority element.
	The majority element is the element that appears more than (n/2) times. You may assume that the majority element alwaus exists in the array.

	Example 1:
	Input: [3,2,3]
	Output: 3

	Example 2:
	Input: [2,2,1,1,1,2,2]
	Output: 2
*/

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 8, 2}, 2},
		{[]int{1, 9, 3, 1, 9, 9}, 9},
		{[]int{1, 1, 2, 1, 1, 9, 9}, 1},
		{[]int{1, 5, 1, 3, 1, 1, 1, 2, 7}, 1},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			// got := majorityElement(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Linear time, O(n) space beacuse the space for the hash table grows with the number of input
func MajorityElement(nums []int) int {
	var result int
	var count int // Store the highest number of appearence for the end loop
	var hash = make(map[int]int)

	// Store the count (number of appreance) of each element
	for _, v := range nums {
		hash[v]++
	}

	// Loop the map, return the key with the most number of counts
	for k, e := range hash {
		if e > count {
			count = e
			result = k
		}
	}

	return result
}
