package main

import (
	"fmt"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

// we need to define a custom type instead of using the raw integer slice
// since we need to define methods on the type to implement the heap interface
type intHeap []int

// Len is the number of elements in the collection.
func (h intHeap) Len() int {
	return len(h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
// This will determine whether the heap is a min heap or a max heap
func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop are used to append and remove the last element of the slice
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
	fmt.Println("oush : ", h)
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// HEAP IMPLEMENTATION
	// g := &intHeap{4, 4, 4, 3, 2, 3}
	// heap.Init(g)

	// // heap.Pop(g)
	// fmt.Println("after pop: ", g)
	// fmt.Println(g)
	// getTopKEl(*g, 2)

	// O(N) IMPLEMENTATION
	nums := []int{1, 4, 2, 4, 3, 1, 2, 4}
	res := getTopKEl(nums, 2)
	fmt.Println("Final result: ", res)
}

func getTopKElHeap(nums []int, k int) []int {
	hMap := map[int]int{}
	for _, v := range nums {
		if _, ok := hMap[v]; !ok {
			hMap[v] = 1
			continue
		}
		hMap[v]++
	}
	return nums
}

func getTopKEl(nums []int, k int) []int {
	res := []int{}
	size := len(nums)
	count := make(map[int]int)
	position := make([][]int, size)

	// For cases where the input array is of len(1)
	if size == 1 && k == 1 {
		return nums
	}

	// Store the number of times the value is repeated in a map
	for _, v := range nums {
		count[v]++
	}

	// loop the count map, place each value in the inner array mapping its count to the position[index]
	for k, v := range count {
		position[v] = append(position[v], k)
	}
	fmt.Println("After sorting into its count position: ", position)

	// Since we are looking for the TOP K elements, we can start from the last index (meaning that these are the values that are repeated the most)
	for i := len(position) - 1; i >= 0; i-- {
		// If there are more than 0 elements in the inner slice, we append them all the the result array
		if len(position[i]) > 0 {
			res = append(res, position[i]...)
		}
		// Break when the result array has K number of values in there
		if len(res) >= k {
			break
		}
	}
	return res
}
