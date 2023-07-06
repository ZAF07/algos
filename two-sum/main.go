package twosum

import "fmt"

/*
  Find the two numbers that when added together, makes the target value
  Example: Target = 9, Args = [2,4,6,7,8]
  Answer: [0, 1] Indexes of 2 & 7 (2 + 7 = 9)

  NOTES: Hash performance is always faster albeit slightly
*/

func main() {
	nums := []int{1, 2, 2, 4, 5}
	res := TwoSum(nums, 7)
	fmt.Println("Result: ", res)
}

func TwoSum(nums []int, target int) []int {
	pair := make(map[int]int) // Data structure to store integer needed to add to equal target

	for k, v := range nums {
		if val, ok := pair[v]; ok {
			return []int{k, val}
		}
		pair[target-v] = k
	}
	return []int{}
}
