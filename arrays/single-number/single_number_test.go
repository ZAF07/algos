package singlenumber

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:
Input: nums = [2,2,1]
Output: 1

Example 2:
Input: nums = [4,1,2,1,2]
Output: 4

Example 3:
Input: nums = [1]
Output: 1

Constraints:
1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears only once.
*/

/* 💡 EXPLANATION

This uses the XOR binary operator (Exclusive OR operator, denoted as ^)

The way binary XOR operator works is that it performs a Bitwise operation on individual bits of two given numbers.
(Bitwise is a level of operation that involves working with individual bits which are the smallest units of data in a computing system.)

Its primary purpose is to compare the corresponding bits of two numbers and return a result where each bit in the output is set to 1 if the two bits being compared are different, and 0 if they are the same.

Here's a truth table that shows how the XOR operator works:
--------------
A	B	A XOR B
0	0	   0
0	1	   1
1	0	   1
1	1	   0
--------------

So given 1 ^ 2 (which translates to 0001 ^ 0010 in bits representation)
The result would be 0011 which is 3 in its decimal notation

0001 (1)
^
0010 (2)
----
0011 (3)


To be really honest, my understanding of the bitwise operation is still very basic as I haven't use it at all in a real world implementations at work or on personal projects.
But here is a visual of finding the single number in this array: [4,1,2,1,2]

1st
0000 (0) Zero here becasue we initialise result at 0. so result(0) ^ n(4)
^
0100 (4)
----
0100 (4)

2nd
0100 (4)
^
0001 (1)
----
0101 (5)

3rd
0101 (5)
^
0010 (2)
----
0111 (7)

4th
0111 (7)
^
0001 (1)
----
0110 (6)

5th & final
0110 (6)
^
0010 (2)
----
0100 (4)
*/

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{9, 3, 1, 9, 1, 4, 4}, 3},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given : %v, Want: %v", tt.nums, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := singleNumber(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}

func singleNumber(nums []int) int {
	var res int
	for _, n := range nums {
		res ^= n
	}
	return res
}
