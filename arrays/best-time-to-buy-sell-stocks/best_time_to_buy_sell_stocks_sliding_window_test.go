package besttimetobuysellstocks

import (
	"fmt"
	"math"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestBestProfitSlidingWindow(t *testing.T) {
	tests := []struct {
		stocks []int
		want   int
	}{
		{[]int{2, 1, 4, 2, 6}, 5},
		{[]int{2, 1, 4, 0, 8}, 8},
		{[]int{0, 1, 4, 2, 6}, 6},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v, Want: %v", tt.stocks, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := bestProfit(tt.stocks)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Sliding window
func bestProfit(stocks []int) int {
	// Start off at first two elements of the array
	buy, sell := 0, 1
	profit := 0
	// As long as the sell position (right index) does not exceed the length of the stocks (array) we keep calculating profit/loss
	for sell < len(stocks) {
		// If buy (left index value) is smaller than the sell value (right index value), that means we make a profit. Calculate the current profit and compare the previous profit
		if stocks[buy] < stocks[sell] {
			sum := stocks[sell] - stocks[buy]
			profit = int(math.Max(float64(profit), float64(sum)))
			fmt.Println(sum, profit, buy, sell)
		} else {
			// If buy is larger than the next sell (sell is smaller than buy), we set the buy value to the sell value
			buy = sell
		}
		// And finally we increase the sell index by 1 for the next comparison
		sell += 1
	}
	return int(profit)
}
