// You are given an array of length n which was originally sorted in ascending order. It has now been rotated between 1 and n times. For example, the array nums = [1,2,3,4,5,6] might become:

// [3,4,5,6,1,2] if it was rotated 4 times.
// [1,2,3,4,5,6] if it was rotated 6 times.
// Given the rotated sorted array nums and an integer target, return the index of target within nums, or -1 if it is not present.

// You may assume all elements in the sorted rotated array nums are unique,

// A solution that runs in O(n) time is trivial, can you write an algorithm that runs in O(log n) time?

// Example 1:
// input: nums = [3,4,5,6,1,2], target = 1
// Output: 4

// Example 2:
// input: nums = [3,5,6,0,1,2], target = 4
// Output: -1

// Constraints:

// 1 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -1000 <= target <= 1000

package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > nums[right]: // unsorted part
			if target < nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		default:
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{3, 4, 5, 6, 1, 2}, 1, 4},
		{[]int{3, 5, 6, 0, 1, 2}, 4, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 0, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	}

	for _, testCase := range testCases {
		result := search(testCase.nums, testCase.target)

		if result != testCase.output {
			fmt.Printf("Test failed: %v and %v , Expected: %v but Got: %v\n", testCase.nums, testCase.target, testCase.output, result)
		} else {
			fmt.Printf("Test passed: %v and %v\n", testCase.nums, testCase.target)
		}
	}
}
