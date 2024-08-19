// Given an integer array nums,
// return an array output where output[i] is the product of all the elements of nums except nums[i].

// Each product is guaranteed to fit in a 32-bit integer.

// Follow-up: Could you solve it in
// O(n) time without using the division operation?

// Input: nums = [1,2,4,6]

// Output: [48,24,12,8]

// Input: nums = [-1,0,1,2,3]

// Output: [0,-6,0,0,0]

// Constraints:

// 2 <= nums.length <= 1000
// -20 <= nums[i] <= 20

package main

import (
	"fmt"
	"reflect"
)

func productExceptSelf(nums []int) {
	leftNums := make([]int, len(nums))
	rightNums := make([]int, len(nums))

	leftNums[0] = 1
	for i := 1; i < len(nums); i++ {
		leftNums[i] = leftNums[i-1] * nums[i-1]
	}

	rightNums[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightNums[i] = rightNums[i+1] * nums[i+1]
	}

	for index, num := range leftNums {
		nums[index] = num * rightNums[index]
	}
}

func main() {
	testCases := []struct {
		nums   []int
		result []int
	}{
		{[]int{1, 2, 4, 6}, []int{48, 24, 12, 8}},
		{[]int{-1, 0, 1, 2, 3}, []int{0, -6, 0, 0, 0}},
	}

	for _, testCase := range testCases {
		productExceptSelf(testCase.nums)

		if !reflect.DeepEqual(testCase.nums, testCase.result) {
			fmt.Printf("Test failed, Expected: %v, Got: %v\n", testCase.result, testCase.nums)
		} else {
			fmt.Println("Test passed")
		}
	}
}
