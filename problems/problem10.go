// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// where nums[i] + nums[j] + nums[k] == 0, and the indices i, j and k are all distinct.

// The output should not contain any duplicate triplets. You may return the output and the triplets in any order.

// Input: nums = [-1,0,1,2,-1,-4]

// Output: [[-1,-1,2],[-1,0,1]]

// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].

// Input: nums = [0,1,1]

// Output: []

// Input: nums = [0,0,0]

// Output: [[0,0,0]]

// 3 <= nums.length <= 1000
// -10^5 <= nums[i] <= 10^5

package main

import (
	"fmt"
	"reflect"
	"slices"
)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	slices.Sort(nums)

	for index, num := range nums {
		if num > 0 {
			break
		}

		if index > 0 && num == nums[index-1] {
			continue
		}

		//two sums
		left := index + 1
		right := len(nums) - 1

		for left < right {
			sum := num + nums[left] + nums[right] // a + b + c

			switch {
			case sum > 0:
				right--
			case sum < 0:
				left++
			default:
				triplet := []int{num, nums[left], nums[right]}
				result = append(result, triplet)
				left++
				right--
				// as duplicate can still exists
				// [-2,0,0,2,2]
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return result
}

func main() {
	testCases := []struct {
		nums   []int
		result [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{[]int{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{[]int{-2, 0, 2}}},
	}

	for _, testCase := range testCases {
		res := threeSum(testCase.nums)

		if !reflect.DeepEqual(res, testCase.result) {
			fmt.Printf("Test failed: %v, Expected: %v but Got: %v\n", testCase.nums, testCase.result, res)
		} else {
			fmt.Printf("Test passed: %v\n", testCase.nums)
		}
	}
}
