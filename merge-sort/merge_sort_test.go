package mergesort

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Example 1:
Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).

Example 2:
Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.

Constraints:
1 <= nums.length <= 5 * 104
-5 * 104 <= nums[i] <= 5 * 104

NOTES: Merge sort has a runtime Big0 of O(nlogn). On larger arrays, Merge sort will be faster than Bubble Sort with a runtiem Big0 of O(n2)
	On smaller arrays though, because Merge Sort has to merge/create new arrays, it might be better to perform Bubble Sort
*/

func TestMergeSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 3, 4, 1, 2}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v, Want: %v", tt.nums, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := merge(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

/*
Example:

	🚨 To be continued. Will try to draw a diagram for better understanding as recursive algos are quite a lot for me to keep track as well!
	1. When merge([2,3]) is called, it will then again call merge, this time with merge(2) & merge(3). This will add to the call stack like so:
	   M[3,2,5,4]
	   ^      ^
	   / \    / \
	   S(3,2)  S(5,4)
	   M[3,2]  M[5,4]
	   ^	    ^
	   /   \     /  \
	   M[2] M[3] M[5] M[4]

2.
*/
func merge(nums []int) []int {
	//  Base case for recursive function
	// We basically keep halving the given array until we get a single element in that array
	// The single element will be returned to the calling stack which will eventually call sort on the returned element which takes in an array and returns an array as well to the calling stack
	if len(nums) < 2 {
		return nums
	}

	// We have to find the midpoint of the array because we want to divide the array into smaller chunks, compare and swap values if needed (smaller or larger, based on your rules)
	// We will end the recursion with two sorted arrays passed to the sort() method
	mid := len(nums) / 2

	// Recursively call the merge() method to further break down the array into smaller chunks, compare and swap values if needed (smaller or larger, based on your rules)
	return sort(merge(nums[mid:]), merge(nums[:mid]))
}

// Function to sort two arrays
func sort(left, right []int) []int {
	result := []int{}
	l, r := 0, 0
	length := len(left) + len(right)
	// You want to loop until the result array is the same lenght as the two given arrays combined because we are sorting them into one array as a result
	for len(result) < length {
		// If left element is smaller than or equal to the right, we simply append the smaller element (left in this case)
		if left[l] <= right[r] {
			result = append(result, left[l])
			// Increase the value of the index/pointer of the smaller element
			l++
		} else {
			// Else we append the right element as it is smaller than the left element
			result = append(result, right[r])
			// Increase the value of the index/pointer of the smaller element
			r++
		}
		// If any of the two index/pointer is == to its array size, we know that anything that comes after is going to be larger, so just append the remaining elements to the result array
		if l == len(left) {
			result = append(result, right[r:]...)
			break
		}
		if r == len(right) {
			result = append(result, left[l:]...)
			break
		}
	}
	return result
}
