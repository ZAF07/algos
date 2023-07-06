package palindrome

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	An integer is considered a palindrome if its sequence reads the same backwards as forwards

	Example 1:
	Input: 121
	Output: true
	Solution: The above is true becaue it reads 121 starting from both left and right

	Example 2:
	Input: 132
	Output: false
	Solution: The above is false becaue it does not read the same starting from both left and right
*/

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		num  int
		want bool
	}{
		{121, true},
		{222, true},
		{11312, false},
		{909, true},
		{11311, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Given : %v, Want: %v", tt.num, tt.want)
		t.Run(testname, func(t *testing.T) {
			res := CheckPalindrome(tt.num)
			t.Logf("Returned from CheckPalindrome(): %v", res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func CheckPalindrome(num int) bool {
	fmt.Println("Given number: ", num)
	r := 0
	for num > r {
		// Get the right most (last digit to the integer) integer of the given number
		r = r*10 + num%10 // times 10 to r to make it 1 tenth number, because if we don't, r would only increase by given number (example: 4 + 1 = 5 or 1 + 2 = 3). We want r to equal the right end of the given number (example: Given number: 121, r: 1 on first pass, and r: 12 on second pass)
		num /= 10         // remove the last integer from the given number for the next pass
		fmt.Println("r: ", r)
		fmt.Println("num: ", num)
	}
	return num == r/10 || num == r
}
