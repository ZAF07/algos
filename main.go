package main

import "fmt"

// Package main is for quick testing of written code without the Go Test package

func main() {
<<<<<<< HEAD
	// nums := []int{1, 5, 2, 4, 3}
	// target := 9
	// res := TwoSum(nums, target)
	// fmt.Println(res)
=======
	nums := []int{0, 3, 4, 3, 4, 5, 1, 6}
	target := 6
	res := TwoSumAdjust2(nums, target)
	fmt.Println(res)
>>>>>>> 8f58bf2 (added two sum with array of numbers (not index) as result)

	// bub := Bubble(nums)
	// fmt.Println(bub)

	// bNums := []int{1, 2, 4, 5, 7, 8}
	// bTarget := 2
	// bin := BinaryS(bNums, bTarget)
	// fmt.Println(bin)
<<<<<<< HEAD

	// nums := []int{4, 5, 6, 0, 1, 2}
	// res := rBinary(nums, 0)
	// fmt.Println(res)

	rNums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4}
	resr := removeDups(rNums)
	fmt.Println("REMOVE DUPS: ", resr)
=======
>>>>>>> 8f58bf2 (added two sum with array of numbers (not index) as result)
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

// THIS IS THE CORRECT ONE SO FAR
func TwoSumAdjust(nums []int, target int) [][]int {
	res := [][]int{}
	seenMap := make(map[int]int)

	for i, v := range nums {
		missingNum := target - v
		if val, seen := seenMap[v]; seen {
			if val == -1 {
				continue
			}
			res = append(res, []int{v, nums[val]})
			seenMap[v] = -1
			continue
		} else {

			if _, ok := seenMap[missingNum]; ok {
				continue
			}
			seenMap[target-v] = i
		}
	}
	return res
}

// 4:t
// v
// 3,  4 , 3  ,4
func TwoSumAdjust2(nums []int, target int) [][]int {
	res := [][]int{}
	seenMap := make(map[int]bool)

	for _, num := range nums {
		missingNum := target - num

		// Check if the missingNum is already in the seenMap,
		// and ensure that the pair doesn't contain duplicate elements.
		if seenMap[num] {
			if !seenMap[missingNum] {
				pair := []int{num, missingNum}
				res = append(res, pair)
			}
		}
		seenMap[missingNum] = true
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

func rBinary(nums []int, target int) int {
	pivot := findPivot(nums)

	if target > nums[0] {
		return rbinary(nums, 0, pivot-1, target)
	} else {
		return rbinary(nums, pivot, len(nums)-1, target)
	}
}

func rbinary(nums []int, left, right, target int) int {
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(nums[mid], target)

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
	return -100
}

//        lr
// [4,5,6,0,1,2]
func findPivot(nums []int) int {
	left, right := 0, len(nums)-1
	ref := nums[0]
	for left < right {
		mid := (left + right) / 2

		if nums[mid] >= ref {
			left = mid + 1
		}

		if nums[mid] < ref {
			right = mid
		}
	}
	return left
}

// map 4:10 3: 9>8>7 2:6 1:4>3 0:1
// index : 0
//                  v
// [0,0,1,1,1,2,2,3,3,3,4] 10
//                v
// [0 0 1 1 1 2 2 3 3 4] 9
//            v
// [0 0 1 1 1 2 2 3 4]
//          v
// [0 0 1 1 1 2 3 4]
func removeDups(nums []int) []int {
	// nums = append(nums[:9], nums[10:]...)
	// fmt.Println(nums)

	// nums = append(nums[:8], nums[9:]...)
	// fmt.Println(nums)

	// nums = append(nums[:6], nums[7:]...)
	// fmt.Println(nums)

	// nums = append(nums[:4], nums[5:]...)
	// fmt.Println(nums)

	// nums = append(nums[:3], nums[4:]...)
	// fmt.Println(nums)

	// nums = append(nums[:1], nums[2:]...)
	// fmt.Println(nums)
	seenMap := make(map[int]int)

	for i := len(nums) - 1; i >= 0; i-- {
		if position, seen := seenMap[nums[i]]; seen {
			nums = append(nums[:position], nums[position+1:]...)
			// seenMap[nums[i]] = i
		}

		seenMap[nums[i]] = i

	}
	return nums
}
