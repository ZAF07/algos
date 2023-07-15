package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
	nums := []int{1, 5, 2, 4, 3}
	target := 9
	res := TwoSum(nums, target)
	fmt.Println(res)

	bub := Bubble(nums)
	fmt.Println(bub)

	bNums := []int{1, 2, 4, 5, 7, 8}
	bTarget := 2
	bin := BinaryS(bNums, bTarget)
	fmt.Println(bin)
}

func TwoSum(nums []int, target int) []int {
	res := []int{}
	pair := make(map[int]int)
	for i, v := range nums {
		if val, hasPair := pair[v]; hasPair {
			res = append(res, val, i)
			return res
		}

		pair[target-nums[i]] = i
	}
	return res
}

func Bubble(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}

func BinaryS(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		}

		if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
