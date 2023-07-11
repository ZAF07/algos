package slidingwindow

/*
Given an array of positive integers nums and a positive integer target, return the minimal length of a
subarray whose sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Example 1:
Input: target = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
Input: target = 4, nums = [1,4,4]
Output: 1

Example 3:
Input: target = 11, nums = [1,1,1,1,1,1,1,1]
Output: 0
*/

// test
func minimumSizeSubarray(nums []int, target int) int {
	min := 0
	window := 1

	for window < len(nums)-1 {
		currentMin := getSum(nums[0 : 0+window]...)
		if currentMin == target && window < min {
			min = window
		}

		for i := 1; i <= len(nums)-window; i++ {
			tempMin := getSum(nums[i : i+window]...)
			if tempMin == target && window < min {
				min = window
			}
		}
	}

	// Check the sum of the entire array at the last step because we are searching for the MINIMAL LENGTH of subarray
	if getSum(nums...) == target {
		return len(nums)
	}

	return min
}
