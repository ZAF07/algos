package main

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/*
	Find the best time to buy & sell stocks to get the max profit
	Example: [2,1,4,5,3] == 4 (we buy at 1 & sell at 5 to make a $4 profit)

*/

func TestCalcBestProfit(t *testing.T) {
	tests := []struct {
		stocks []int
		want   int
	}{
		{[]int{2, 1, 4, 5, 3}, 4},
		{[]int{2, 13, 4, 5, 3}, 11},
		{[]int{2, 1, 421, 51, 35}, 420},
		{[]int{2, 10, 153, 5, 1}, 151},
		{[]int{2, 11, 4, 2, 9}, 9},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.stocks, tt.want)
		t.Run(testName, func(t *testing.T) {
			got := CalcBestProfit(tt.stocks)
			assert.Equal(t, tt.want, got)
		})
	}

}

func CalcBestProfit(stocks []int) int {
	minPrice := 1000
	maxProfit := 0
	for _, s := range stocks {
		// for each pass, check if current value is < min_price (this is to find the smallest value first). if so, set current value as minPrice
		if s < minPrice { // Finding the buying value first. That's why we continue the loop from here
			minPrice = s
			continue
		}

		// On subsequent pass(after finding the smallest value so far. this means we are now finding the sell value), we want to check if current value - minPrice is > maxProfit (which is set to 0 by default) and if so, we set the sum as the maxPrice
		if s-minPrice > maxProfit {
			maxProfit = s - minPrice
		}
	}
	return maxProfit
}

// Works the same as above but has extra work to perform with the omitted continue keyword (Above method should be optimal) o(N)
func myOwn(stocks []int) int {
	low := 100
	high := 0

	for _, v := range stocks {
		if v < low {
			low = v
		}

		if v-low > high {
			high = v - low
		}
	}
	return high
}
