// Given an array of integers nums and an integer target,
// return the indices i and j such that nums[i] + nums[j] == target and i != j.

// Input:
// nums = [3,4,5,6], target = 7

// Output: [0,1]

// Input: nums = [4,5,6], target = 10

// Output: [0,2]

// Input: nums = [5,5], target = 10

// Output: [0,1]

package main

import "fmt"

func twoSum(nums []int, target int) (int, int) {
	hmap := make(map[int]int)

	for index, num := range nums {
		x := target - num

		if xIndex, ok := hmap[x]; ok {
			return xIndex, index
		}

		hmap[num] = index
	}

	return -1, -1
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		i      int
		j      int
	}{
		{[]int{3, 4, 5, 6}, 7, 0, 1},
		{[]int{4, 5, 6}, 10, 0, 2},
		{[]int{5, 5}, 10, 0, 1},
	}

	for _, testCase := range testCases {
		i, j := twoSum(testCase.nums, testCase.target)
		if i != testCase.i && j != testCase.j {
			fmt.Printf("Test Case: %v failed, Expected: %v and %v, Got: %v and %v\n", testCase.nums, testCase.i, testCase.j, i, j)
		} else {
			fmt.Printf("Test Case: %v passed\n", testCase.nums)
		}
	}
}
