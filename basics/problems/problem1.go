package main

import "strings"


// ##
// ##	dyanmic programming 
// ##
func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}

func reverseNumber(num int) int {
	//neutralizing a negative number to positive number
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	result := 0

	for num > 0 {
		last := num % 10
		result = result*10 + last
		num = num / 10
	}

	if isNegative {
		result = -result
	}

	return result
}



// ##
// ##	Two pointers problems
// ##
func reverseString(str string) string {
	// why this solution as go string is immutable
	// and proper UTF-8 support
	var result strings.Builder

	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		result.WriteRune(runes[i])
	}

	return result.String()
}

func validPalindrome(str string) bool {
	// the solution is case sensitive
	// UTF-8 encoding support
	runes := []rune(str)

	for i := 0; i < len(runes)-1; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}

func largestElement(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}

// ##
// ##	hash set or hash map problems
// ##
func removeDuplicate(nums []int) int {
	hset := make(map[int]struct{})

	i := 0

	for _, num := range nums {
		if _, ok := hset[num]; !ok {
			hset[num] = struct{}{}
			nums[i] = num
			i++
		}
	}

	return i
}

// ##
// ##	bitwise manipulation problmes
// ##
func evenOrOdd(num int) string {
	if num&1 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func missingElement(nums []int) int {
	// nums contains range from [0 to n]
	// all elements are distinct element
	// nums = []int{ 0, 2}
	// ans = 1
	// 2 ^ 2 = 0
	// 2 ^ 0 = 2
	// 2 ^ 3 = 3 ^ 2

	allXored := 0
	for i := 1; i < len(nums)+1; i++ {
		allXored ^= i
	}

	for _, num := range nums {
		allXored ^= num
	}

	return allXored
}

func singleNumber(nums []int) int {
	// Given a non-empty array of integers nums, 
	// every element appears twice except for one. Find that single one
	// 2 ^ 2 = 0

	target := nums[0]

	for i := 1 ; i < len(nums); i++ {
		target ^= nums[i]
	}

	return target
}