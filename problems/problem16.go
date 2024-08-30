// You are given an array of distinct integers nums, sorted in ascending order, and an integer target.

// Implement a function to search for target within nums. If it exists, then return its index, otherwise, return -1.

// Your solution must run in O(log n) time.

// Example 1:
// input: nums = [-1,0,2,4,6,8], target = 4
// Output: 3

// Example 2:
// input: nums = [-1,0,2,4,6,8], target = 3
// Output: -1

// Constraints:

// 1 <= nums.length <= 10000.
// -10000 < nums[i], target < 10000

package main

import "fmt"

func bs(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		default:
			return mid
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		input  []int
		target int
		output int
	}{
		{[]int{-1, 0, 2, 4, 6, 8}, 4, 3},
		{[]int{-1, 0, 2, 4, 6, 8}, 3, -1},
		{[]int{-1, 0, 2, 4, 6, 8}, -1, 0},
	}

	for _, testCase := range testCases {
		position := bs(testCase.input, testCase.target)

		if position != testCase.output {
			fmt.Printf("Test failed: %v, Expected: %v but Got: %v\n", testCase.input, testCase.output, position)
		} else {
			fmt.Printf("Test passed: %v\n", testCase.input)
		}
	}
}
