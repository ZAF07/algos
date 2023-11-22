package balloon

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
 Taken from Codility
	Given a string, find how many times you can remove letters from the string to form a new word ("BALLOON"). Each time you remove letters from the given string, you can only remove 7 letters. If there is still enough letters left, the word "BALLOON" may still be found in the given string.
*/

func TestBalloon(t *testing.T) {
	tests := []struct {
		given string
		want  int
	}{
		{"hello", 0},
		{"LOLOANACD", 0},
		{"OBGAIKSROLWLN", 1},
		{"OABLGAINBKSLOOROLWLN", 2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v\n", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := balloon(tt.given)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

func balloon(str string) int {
	if len(str) < 7 {
		return 0
	}

	var res int

	count := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		if _, seen := count[str[i]]; !seen {
			count[str[i]] = 1
		} else {
			count[str[i]]++
		}
	}

	res = count['B']
	res = int(math.Min(float64(count['A']), float64(res)))
	res = int(math.Min(float64(count['L']/2), float64(res)))
	res = int(math.Min(float64(count['O']/2), float64(res)))
	res = int(math.Min(float64(count['N']), float64(res)))

	return res
}
