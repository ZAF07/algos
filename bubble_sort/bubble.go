package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 2, 9, 6, 1, 8}
	result := bubble(nums)
	res := bubbleTwo(nums)
	fmt.Println(result)
	fmt.Println(res)
}

func bubble(arr []int) []int {
	// outer pass goes through the entire length of the array
	for i := 0; i < len(arr); i++ {
		// inner pass goes through the entire length of the array as well
		for j := 0; j < len(arr)-1-i; j++ { // We always decrease the inner loop by i so that we dont visit the last element in each pass. Because we already know that after each pass, the largest element will be pushed to the back of the slice
			// check if current array element is larger then the element right next to it
			if arr[j] > arr[j+1] {
				// if so, swap the current element with the element right next to it
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func bubbleTwo(nums []int) []int {
	// compare current and next value, if next is > current, swap positions
	for i := 0; i < len(nums); i++ { // First pass goes through the entire array.
		for j := 0; j > len(nums)-1-i; j++ { // inner loop keeps decreasing because we know that at each pass, the last element will be the largest element
			if nums[j] > nums[j+1] { // if the current num is > next num, we swap positions
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
