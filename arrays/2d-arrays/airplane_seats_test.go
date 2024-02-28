package darrays

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*

	We are writing an airplane seat allocation algorithm. We are given n=number of rows and reserved= string representing the reserved seats.
	Families can either be seated in a direct row of 4 or can be seperated by one space with two family members at each side.
	The seats in each row are placed as follow: [A B C  D E F G  H J K]
	Return the maximum number of famiy of four we can sit together given the input.

	Example:
	Given rows = 1, reserved = "1A 1k"
	Result: 2. Each family is seperated by one space with 2 family mambers side by side [x B C  D E f g  h j x]
*/

func TestAirplaneSeats(t *testing.T) {
	tests := []struct {
		n        int
		reserved string
		want     int
	}{
		{1, "", 2},                  // 2 families
		{1, "1K", 2},                // 2 families
		{1, "1A 1K", 2},             // 2 families
		{1, "1A", 2},                // 2 families
		{1, "1A 1F 1K", 1},          // left only
		{1, "1A 1F 1K 1G", 1},       // left only
		{1, "1A 1K 1B", 1},          // right only
		{1, "1A 1C 1B 1D", 1},       // right only
		{1, "1A 1C 1B 1H 1J 1K", 1}, // middle only
		{1, "1A 1B 1H 1J 1K", 1},    // middle only
		{1, "1A 1H 1K", 1},          // middle only
		{2, "", 4},                  // 4 families
		{10, "1A 1K", 20},           // 2 families
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v & %v || Want: %v", tt.n, tt.reserved, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := calculateFamiliesInRow(tt.n, tt.reserved)
			want := tt.want
			assert.Equal(t, want, got)
		})
	}
}

func calculateFamiliesInRow(n int, reserved string) int {
	// If reserved seat is empty, return 2 (each row can sit a maximum of 2 families)
	if reserved == "" {
		return 2 * n
	}
	var total int

	// 2d array representing rows and cols
	rows := make([][]bool, n)

	// init the inner array to type bool and size = 10
	for r := range rows {
		rows[r] = make([]bool, 10)
	}

	// create a regexp to capture the row and seat ID based on the given string
	seatRowPattern := regexp.MustCompile(`(\d+)([A-Za-z]+)`)

	// iterate the split instance of the given reserved string and pass it to the regexp
	for _, r := range strings.Split(reserved, " ") {
		matched := seatRowPattern.FindStringSubmatch(r)

		// get the offset unicode value of the row and seat by performing -'A' for alphabet || -'0' for integers
		row, seat := int([]rune(matched[1])[0]-'1'), int([]rune(matched[2])[0]-'A')

		// because I is skipped in the seat sequence, we have to perform the check and subtract K by 1 position
		if matched[2] == "K" {
			seat--
		}
		// set the row and seat position to true in the 2d array representing the seats reserved
		rows[row][seat] = true
	}

	// iterate the row and cols and calculate how many family of 4 can fit into each row given the reserved seats
	for _, row := range rows {
		// left & right with no space in the middle
		if canSit(row, 1, 4) && canSit(row, 5, 8) {
			total += 2
			fmt.Println("✅ 2 families")
			continue
		}

		// middle only & no space for 4 on left and right
		if canSit(row, 3, 6) || canSit(row, 1, 4) || canSit(row, 5, 8) {
			total++
			fmt.Println("✅ left, right or middle")
			continue
		}
	}

	return total
}

// Loop the given row and return false if a seat between sequence from start to end if reserved (true)
func canSit(row []bool, start, end int) bool {
	for i := start; i <= end; i++ {
		if row[i] {
			return false
		}
	}
	return true
}
