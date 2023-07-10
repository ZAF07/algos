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
	buy, sell := 0, 1
	profit := 0
	for sell < len(stocks) {
		if stocks[buy] < stocks[sell] {
			sum := stocks[sell] - stocks[buy]
			profit = int(math.Max(float64(profit), float64(sum)))
			fmt.Println(sum, profit, buy, sell)
		} else {
			buy = sell
		}
		sell += 1
	}
	return int(profit)
}
