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

func TestLongestPalindromeOptimised(t *testing.T) {
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
			got := LongestPalindromeManacher(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

/*
Manacher's algprithm 0(n) time complexity

This is a very efficient algoritum for solving the Longest palindrome in a substring proplem
Because we are only iterating the string once and expanding the left and right pointers in constant time, this makes it a very efficient algorithm for this probelem
*/
func LongestPalindromeManacher(s string) string {
	// Edge case for when there is only one character in the string
	if len(s) < 2 {
		return s
	}

	var result string
	// Loop the string only once
	for i := 0; i < len(s)-1; i++ {

		// For odd palindrome (ana, bab), we start at the same point and expand towards the left and right
		l, r := i, i

		// Here we check that l>=0 (so that we don't go past the starting character of the string) and that r<= len(s) (so that we don't go past the last character of the string)
		// Also we are continuing with the inner loop as long as we find that both s[l] && s[r] are the same characters. Tat means theay are in fact palindromes
		// As long as s[l] == s[r], we keep expanding
		for l >= 0 && r < len(s) && s[l] == s[r] {
			// Since we are looking for the longest palindrome substring, we only change the result if the length of the current palindrome is larger than the palindrome in the current result
			if (r-l)+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}

		// For even palindrome (anna, baab), we start at the two middle points (nn,aa) and keep expanding to the left and right
		l, r = i, i+1

		// Make sure not to go past the start and end of the string
		// Continue the loop only if s[l] == s[r]
		// If s[l] == s[r], we check if the length of the current palindrome is larger than the current result (we are looking for the longest palindrome in the substring)
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r-l)+1 > len(result) {
				result = s[l : r+1]
			}
			l--
			r++
		}
	}
	return result
}
