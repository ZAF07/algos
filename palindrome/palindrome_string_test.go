package palindrome

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
A string is considered a palindrome if a string can be read the same both backwards and forwards

Eamplpe 1:
input: "anna"
output: true

Description: The string is "anna" when read forwards and backwards.
*/

func TestPalindromeString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"anna", true},
		{"xanna", false},
		{"babba", false},
		{"xsbabsx", true},
		{"ac", false},
		{"a", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.s, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := PalindromeString(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}

}

func PalindromeString(s string) bool {
	if len(s) < 2 {
		return true
	}

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
