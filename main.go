package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
	nums := []int{1}
	// res := merge_sort(nums, "start")
	// fmt.Println(res)

	nums2 := []int{1, 8}
	r := sort(nums, nums2, "start")
	fmt.Println(r)
}

func merge_sort(nums []int, location string) []int { // >> ms(3,2, 1,5) > ms(3,2) >> ms(3) ms(2) > s([3],[2])  ms(5,1) >> ms(5) ms(1) > s([5].[1])
	fmt.Println("1️⃣: ", nums, location)
	if len(nums) < 2 {
		return nums
	}

	// Get the midpoint of the array
	mid := len(nums) / 2

	fmt.Println("Calling sort with : ", nums[:mid], nums[mid:], location)
	// Recursively sort the halved array
	return sort(merge_sort(nums[:mid], "left"), merge_sort(nums[mid:], "right"), location)
}

// Result: [2, 3, 5, 1]
//
//	l  r
//
// [2,3], [5,1]
func sort(left, right []int, location string) []int {
	fmt.Println("2️⃣✅ In sort: ", left, right, location)
	result := []int{}
	lIndex, rIndex := 0, 0
	// 0 < 4
	for len(result) < len(left)+len(right) {
		// If left element <= right element, append element to result and increase left index
		if left[lIndex] <= right[rIndex] {
			result = append(result, left[lIndex])
			lIndex++
		} else {
			// This means that the right element is < left element. Append right element to result
			result = append(result, right[rIndex])
			rIndex++
		}

		if rIndex == len(right) {
			result = append(result, left[lIndex:]...)
			break
		}

		if lIndex == len(left) {
			result = append(result, right[rIndex:]...)
			break
		}
	}
	fmt.Println("3️⃣👍: ", result)
	return result
}
