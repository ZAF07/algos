package palindrome

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"gopkg.in/go-playground/assert.v1"
)

/* 125 Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{" race a car.", false},
		{" ", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v | Want: %v", tt.input, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := validPalindrome(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func validPalindrome(s string) bool {
	// Pointers start at first and last character of the string
	l, r := 0, len(s)-1

	for l <= r {

		// Check if left is an alpha-numeric character
		for l < r && !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			// Move left pointer up one position
			l++
		}

		// Check if right is an alpha-numeric character
		for l < r && !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			// Move right pointer down one position
			r--
		}

		// Check that both characters at left and right pointers match
		if unicode.IsLetter(rune(s[l])) && unicode.IsLetter(rune(s[r])) {
			equal := strings.EqualFold(string(s[l]), string(s[r]))
			if !equal {
				return false
			}

		} else {
			// Checking for number match
			if s[l] != s[r] {
				return false
			}
		}
		l++
		r--
	}
	return true
}
