package palindrome

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"anna", "anna"},
		{"xanna", "anna"},
		{"babba", "abba"},
		{"xsbabxr", "bab"},
		{"ac", "a"},
		{"a", "a"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.s, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := LongestPalindrome(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}

}

/*
Basic approach. O(n^2) (O of n squared) (See palindrome/longest_palindrome_manacher_test.go for an optimised solution)

Not very optimal because we are looping twice here.
The outer loop loops the string once, while the inner loop does the exact same, checking if the string is a palindrome by expanding to the right only
*/
func LongestPalindrome(s string) string {
	// Edge case
	if len(s) < 2 {
		return s
	}
	var result string
	// Outer loop, Left pointer starts at the current position
	for i := 0; i < len(s); i++ {
		r := i

		// Inner loop, Right pointer starts at the current pointer as well. Inefficient...
		for i >= 0 && r < len(s) {
			currentStrLength := (r - i) + 1
			if isPalindrome(s[i : r+1]) {
				if currentStrLength > len(result) {
					result = s[i : r+1]
				}
			}
			r++
		}
	}
	return result
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
