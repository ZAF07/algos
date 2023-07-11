package topkfrequentnums

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]

Hard: Return the larger of two elements if there are more than 1 elements
Input: nums = [1,1,1,3,3,3,2,2,2,5], k =  2
Output: [2,3] (2 > 1)
*/

func TestTopKFrequentNums(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 2, 1, 3}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 2, 1, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 2, 1, 3, 1}, 1, []int{1}},
		{[]int{1, 2, 2, 1, 3, 3, 3, 1}, 2, []int{1, 3}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want : %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := GetTopKEl(tt.nums, tt.k)
			assert.Equal(t, tt.want, got) // 🚨 Slight bug here. The returned slice may not be in the same order as tt.want. But the answer is still correct. Fix this when time permits
		})
	}
}

func GetTopKEl(nums []int, k int) []int {
	res := []int{}
	size := len(nums)
	count := make(map[int]int)
	position := make([][]int, size)

	// For cases where the input array is of len(1)
	if size == 1 && k == 1 {
		return nums
	}

	// Store the number of times the value is repeated in a map
	for _, v := range nums {
		count[v]++
	}

	// loop the count map, place each value in the inner array mapping its count to the position[index]
	for k, v := range count {
		position[v] = append(position[v], k)
	}
	fmt.Println("After sorting into its count position: ", position)

	// Since we are looking for the TOP K elements, we can start from the last index (meaning that these are the values that are repeated the most)
	for i := len(position) - 1; i >= 0; i-- {
		// If there are more than 0 elements in the inner slice, we append them all the the result array
		if len(position[i]) > 0 {
			res = append(res, position[i]...)
		}
		// Break when the result array has K number of values in there
		if len(res) >= k {
			break
		}
	}
	return res[0:k]
}
