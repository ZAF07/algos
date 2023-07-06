package main

import "fmt"

func main() {
	num := 121
	ispalindrome := checkPalindrome(num)
	fmt.Println(ispalindrome)
	// last := nums % 10 // gives back the remainder of the num / 10
	// lastRemoved := nums / 10 // gives you the rounded up (because this data type is an integer) division of num / 10 without the remainder
}

func checkPalindrome(num int) bool {
	fmt.Println("Given number: ", num)
	r := 0
	for num > r {
		// Get the right most (last digit to the integer) integer of the given number
		r = r*10 + num%10 // times 10 to r to make it 1 tens number, because if we don't, r would only increase by given number (example: 4 + 1 = 5 or 1 + 2 = 3). We want r to equal the right end of the given number (example: Given number: 121, r: 1 on first pass, and r: 12 on second pass)
		num /= 10         // remove the last integer from the given number for the next pass
		fmt.Println("r: ", r)
		fmt.Println("num: ", num)
	}
	return num == r/10 || num == r
}
