package main

import (
	"fmt"
	"math"
)

// Package main is for quick testing of written code without the Go Test package

func main() {

	s := "bascfgtb"
	res := aa(s)
	fmt.Println(res)
}

func maxLenSubarray(s string) int {

	if s == "" {
		return 0
	}
	left, right := 0, 0
	maxLen := 0
	seenMap := make(map[byte]int)

	for right < len(s) {
		// If we've seen this letter & it is somewhere behind the left pointer (meaning that it is not included in the current window) we simply calculate the current window length and expand the window again
		if location, seen := seenMap[s[right]]; !seen {
			seenMap[s[right]] = right
		} else {
			// If last seen letter comes after the current left pointer (meaning that the current window includes the last seen letter on the right pointer AND somewhere in the window) we simply push left pointer to the last seen position and continue sliding the window
			if location >= left {
				// We +1 here because we know that the right pointer of the window is == the seen letter "aaba"
				//  l s   r
				// "a a b a "
				left = location + 1
				continue
			}
			seenMap[s[right]] = right // We update the location of the current seen letter
		}
		subLen := right - left + 1
		maxLen = int(math.Max(float64(subLen), float64(maxLen)))
		right++
	}
	return maxLen
}

func aa(s string) int {
	left, right := 0, 0
	maxLen := 0
	seenMap := make(map[byte]int)

	for right < len(s) {
		if location, seen := seenMap[s[right]]; seen {
			if location >= left {
				left = location + 1
				continue
			}
		}

		seenMap[s[right]] = right // We update the location of the current seen letter
		subLen := right - left + 1
		maxLen = int(math.Max(float64(subLen), float64(maxLen)))
		right++
	}
	return maxLen
}
