package classroomseats

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

// TODO: Make more test cases

/*
Given an array layout, with its first element representing the total number of desks in a classroom and elements [1:] representing the occupied seats,
calculate the numnber of possible seat allocations 2 new students can be allocated to such that they are either next to or across each other.

Example:
Given: [6, 4]
o,o
o,t
o,o
Answer: 4

Given: [8, 2, 5]
o,t
o,o
t,o
o,o
Answer: 5
*/

func TestClassroomSeats(t *testing.T) {
	tests := []struct {
		given []int
		want  int
	}{
		{[]int{6, 4}, 4},
		{[]int{3, 2}, 1},
		{[]int{3, 1}, 0},
		{[]int{8, 2, 5}, 5},
		{[]int{12, 4, 10, 9}, 8},
		{[]int{12, 4, 9}, 10},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given %v || Want: %d", tt.given, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := searAllocations(tt.given)
			want := tt.want
			assert.Equal(t, got, want)
		})
	}
}

/*
o,o
o,t
o,o
o,o
t,t
o,o
*/
func searAllocations(layout []int) int {
	var count int
	size := layout[0]
	deskLayout := make([]bool, size)

	// Populate the desk layout
	for _, i := range layout[1:] {
		deskLayout[i-1] = true
	}
	fmt.Println("Desk layout --> ", deskLayout)
	for c := 0; c < len(deskLayout)-1; c++ {
		// If current seat is already taken, skip the calculations. No one can take an occupied seat
		if deskLayout[c] == true {
			continue
		}
		r := c + 1
		l := c + 2

		// last row. we only check the sear directly to the right of current. Early termination is possible here as there are no more seats to calculate
		if c%2 == 0 && c == size-2 {
			if canSit(deskLayout[c], deskLayout[r]) {
				count++
			}
			break
		}

		// Left side. Check seat directly to the right and behind
		if c%2 == 0 {
			if canSit(deskLayout[c], deskLayout[r]) {
				count++
			}
			if canSit(deskLayout[c], deskLayout[l]) {
				count++
			}
		} else {
			// Right side. Check seat directly behind only (if exists)
			// check that current seat is on the right side (means we need to check 2 seats after current) and that 2 seats after current (Directly behind) exists
			// Current could be the second last seat, so there may not be a seat directly behind current
			if l <= size-1 && canSit(deskLayout[c], deskLayout[l]) {
				count++
			}
		}

	}
	return count
}

func canSit(l, r bool) bool {
	return !l && !r
}
