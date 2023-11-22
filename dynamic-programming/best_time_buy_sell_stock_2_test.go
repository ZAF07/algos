package dynamicprogramming

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

/* 122. Best Time to Buy and Sell Stock II
 You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
 On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.
 Find and return the maximum profit you can achieve.

 Example 1:
 Input: prices = [7,1,5,3,6,4]
 Output: 7
 Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
 Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
 Total profit is 4 + 3 = 7.

 Example 2:
 Input: prices = [1,2,3,4,5]
 Output: 4
 Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
 Total profit is 4.

 Example 3:
 Input: prices = [7,6,4,3,1]
 Output: 0
 Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

Constraints:
 1 <= prices.length <= 3 * 104
 0 <= prices[i] <= 104
*/

/*
	Solution in text

	First we init two DS to store both the values of the sell and buy
	EG. When we buy today, how much would we earn and same for the sell

	before the loop, we calculate the profit/balance from buying/selling a stock today

	start a loop starting at index 1 (day 2) and compare the profit from buying/selling today vs the previous day

	At the end of the loop we would have picked the maximum value over the days
*/

// TODO: FIND THE CORRECT PROBLEM STATEMENT FOR THIS PROBLEM
func TestBestTimeToBuyAndSellStocks(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{17, 6, 4, 1, 13, 1, 9, 10}, 21},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("Given: %v || Want: %v\n", tt.prices, tt.want)

		t.Run(testName, func(t *testing.T) {
			got := maxProfit(tt.prices)
			want := tt.want
			assert.Equal(t, got, want)
		})
	}
}

// 7,1,4,5,2,10
func maxProfit(prices []int) int {
	stocks := len(prices)
	if stocks <= 1 {
		return 0
	}

	buy := make([]int, stocks)  // -7, -1, -1, -1, 2, 10
	sell := make([]int, stocks) // 0, 0, 3, 4, 4, 12

	// Buying on the first day, so our balance will be negative
	buy[0] = -prices[0]

	// Selling (more like holding) on the first day. We don't own any stocks so there is 0 profit
	sell[0] = 0

	// Starts on the second day, as we already calculated the p&l for the first day above
	for i := 1; i < stocks; i++ {
		// Should i keep what i bought yesterday? or should i buy today (to buy, i have to take what i earned yesterday and deduct that by today's price)
		buy[i] = max(buy[i-1], sell[i-1]-prices[i]) // hold what i bought yesterday or sell and buy today

		// Should i hold what i earned yesterday or should i sell today? (to figure this out, i should see if whatever i bought yesterday + today's prices is larger than what i had yesterday)
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]) // hold (remain as per yesterday) or sell from yesterdays purchase
	}

	return sell[stocks-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
