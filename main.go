package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
	nums := []int{1} // len = 6
	fmt.Println(search(nums, 0))
}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	pivot := findPivot(nums)

	if pivot == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if target < nums[0] {
		return binarySearch(nums, pivot, len(nums)-1, target)
	} else {
		return binarySearch(nums, 0, pivot-1, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	// 5,4
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// Find the pivot index if it exists, return -1 otherwise
func findPivot(nums []int) int {
	// Since all elements in the array are unique,
	// if the array is rotated, the first element will be greater than the last element.
	// Hence we can do an early return if the first element is less than the last element.
	if nums[0] < nums[len(nums)-1] {
		return -1
	}

	// Binary search for the pivot index
	// The element at the pivot index will be the largest element in the array
	ref := nums[0]

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		//💡WHY DO WE BRING THE LEFT POINTER UP IF NUM[MID] >= NUMS[0]??
		// Because we know that the given array is already sorted in ascending order and we are trying to find the smallest possible element in the given array, when a rotation happens within an ordered array, the smallest element (the first element in an unsorted array) will be moved to the right k number of times. So if the middle element is larget than the first element in a rotated array, we know that the smallest possible element will definitely be somewhere to the right
		if nums[mid] >= ref {
			left = mid + 1
		}
		// 💡WHY DO WE BRING THE RIGHT POINTER DOWN IF NUMS[MID] < NUMS[0] ??
		// Because we are trying to find the smallest possible element here (the pivot or where we can split the rotated array into two to easily perform Binary Search on afterwards) know that the array is already sorted in ascending order, any number to the right of an element will only grow larger (unless the array is rotated of course!)
		// So if the current middle element (nums[(left + right) / 2]) is lesser than the first element (nums[0]), knowing that elements to the right are only going to be larger, there is the possibility that num[mid + 1] could be larger than or equal to nums[0], so we can safely narrow down the search for the smallest element from the current left pointer up until the middle element
		if nums[mid] < ref {
			right = mid
		}
	}

	fmt.Println("THE P[IVOT]: ", left)
	return left
}
