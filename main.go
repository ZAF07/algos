package main

import (
	"fmt"
	"math"
)

// Package main is for quick testing of written code without the Go Test package

func main() {
	nums := []int{1, 2, 3, 4, 2} // len = 6 | max = 15
	// nums := []int{4, 5, 6, 10, 120, 2, 3} // len = 6 | max = 141
	// nums := []int{4, 2} // len = 6 | max = 142
	// nums := []int{2, 4} // len = 6 | max = 142
	fmt.Println(minimumSizeSubarray(nums, 9)) // == 2
}

// THIS IS A DYNAMIC SIZE SLIDING WINDOW
func minimumSizeSubarray(nums []int, target int) int {
	left := 0
	right := 0
	sum := 0
	min := len(nums)
	//    l              r
	// 4, 5, 6, 0, 1, 2, 3
	// loop through the entire array
	for right < len(nums) {
		sum += nums[right]

		//  We have found a subarray with N size that is >= target
		// For example: [2,3,1,5] >= 7
		// But [3,1,5] is also >= 7 and this is a better minimum size subarray
		// So, as long as the sum is >= target, we would want to keep bringing the left pointer to the right (shrink the window from the left),  recalculate the sum of the window until sum < target, and update the minimun len result
		for sum >= target {
			// 💡 IF YOU ARE TYING TO LOOK FOR SUM == TARGET, USE THIS
			// if sum == target {
			// 	size := right - left + 1
			// 	min = int(math.Min(float64(size), float64(min)))
			// }

			// UNCOMMENT IF YOU ARE LOOKING FOR SUM >= TARGET
			size := right - left + 1
			min = int(math.Min(float64(size), float64(min)))

			// Shrink the window from the left side. There could also be a smaller window size where sum >= target
			sum -= nums[left]
			// Dont fotget to bring the left pointer to the right (Shrik the window)
			left++
		}

		// When we get here, this means that we have shrunk the window from the left side in (substracting the sum from the left pointer) until the sum is now < target or the sum is < target
		// We now move the right pointer, growing our sliding window again and recalculating the sum of the window
		// Basically, we shrink when sum >= target and we expand when sum < target
		right++
	}

	// As long as the total sum is >= target, keep substracting total sum from the value at the left pointer and then add one to the left pointer
	// if the subarray's length is < min value, change the min value (shrink the window here to get a potential smaller window)

	// Add one to the right pointer if the window shrinks from the left
	if min == len(nums) {
		return 0
	}
	return min

}
