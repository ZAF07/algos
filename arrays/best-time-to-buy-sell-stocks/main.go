package main

import (
	"fmt"
)

func main() {

	stocks := []int{3, 1, 3, 2, 21}
	res := calcBestProfit(stocks)
	res1 := myOwn(stocks)
	fmt.Println(res, res1)
}

func calcBestProfit(stocks []int) int {
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
