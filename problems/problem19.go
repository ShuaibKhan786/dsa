// You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:

// [3,4,5,6,1,2] if it was rotated 4 times.
// [1,2,3,4,5,6] if it was rotated 6 times.
// Notice that rotating the array 4 times moves the last four elements of the array to the beginning. Rotating the array 6 times produces the original array.

// Assuming all elements in the rotated sorted array nums are unique, return the minimum element of this array.

// A solution that runs in O(n) time is trivial, can you write an algorithm that runs in O(log n) time?

// Example 1:
// input: nums = [3,4,5,6,1,2]
// Output: 1

// Example 2:
// input: nums = [4,5,0,1,2,3]
// Output: 0

// Example 3:
// input: nums = [4,5,6,7]
// Output: 4

// Constraints:

// 1 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000

package main

import "fmt"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	testCases := []struct {
		nums   []int
		output int
	}{
		{[]int{3, 4, 5, 6, 1, 2}, 1},
		{[]int{4, 5, 0, 1, 2, 3}, 0},
		{[]int{4, 5, 6, 7}, 4},
		{[]int{-10, 1, 2, 3, 4, 5, 6, 8}, -10},
	}

	for _, testCase := range testCases {
		result := findMin(testCase.nums)

		if result != testCase.output {
			fmt.Printf("Test failed: %v, Expected: %v but Got: %v\n", testCase.nums, testCase.output, result)
		} else {
			fmt.Printf("Test passed: %v\n", testCase.nums)
		}
	}
}
