package main

import (
	"fmt"
)

/*
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.

EXAMPLE:
Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
*/

func main() {
	nums := []int{1, 9, 9}
	res := plusOne(nums)
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// Because we are starting from the end of the slice, if the element is < 9 (1 - 8), we can simply increment its value by 1 and return the entire slice
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// else we set the current value to 0, because 9 + 1 will == 10,
		digits[i] = 0

	}
	// This will only be run if the remaining last element in the loop (meaning the first element digits[0] since we are going backwards) is 9. If so, we append 1 to the start of the the slice and return it
	digits = append([]int{1}, digits...)
	return digits
}
