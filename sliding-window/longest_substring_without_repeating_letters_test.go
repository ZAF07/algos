package slidingwindow

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
Given a string s, find the length of the longest
substring
 without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func TestLongestSubString(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aabbab!a", 3},
		{"aababcd", 4},
		{"bbcaecbb", 4},
		{"aba", 2},
		{"aa", 1},
		{"a", 1},
		{"", 0},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.s, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := longestSubstring(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func longestSubstring(s string) int {
	// Two pointers representing the size of the current window
	left, right := 0, 0
	maxLen := 0
	// Map to store letters we have seen
	seenMap := make(map[byte]int)

	// Loop the entire string
	for right < len(s) {
		// So if we have seen the letter before, we first check if it is currently before of after the left pointer
		// If it is AFTER or IS the current left pointer, we simply shift the left pointer one letter after the seen letter and continue with the loop
		// We don't have to do anything else here. The right pointer will not increment as well. Only the left pointer in this case. The next iteration will recheck the same conditions and calculate the length of the new window if there are no repeating letters
		if location, seen := seenMap[s[right]]; seen {
			if location >= left {
				left = location + 1
				continue
			}
		}
		// Regardless of whether we have seen the letter or not, we update the seen map with the position of the current letter
		// We then calculate the length of the current window and increment the right pointer by one
		seenMap[s[right]] = right // We update the location of the current seen letter
		subLen := right - left + 1
		maxLen = int(math.Max(float64(subLen), float64(maxLen)))
		right++
	}
	return maxLen
}
