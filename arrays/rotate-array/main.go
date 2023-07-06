package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	res := rotateArr(nums, 2)
	fmt.Println(res)
}

func rotateArr(nums []int, k int) []int {
	x := len(nums) - k%len(nums)

	nums = append(nums[x:], nums[:x]...)
	return nums
}
