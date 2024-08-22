// Given an array of integers numbers that is sorted in non-decreasing order.

// Return the indices (1-indexed) of two numbers, [index1, index2], such that they add up to a given target number target and index1 < index2. Note that index1 and index2 cannot be equal, therefore you may not use the same element twice.

// There will always be exactly one valid solution.

// Your solution must use O(1) additional space.

// Input: numbers = [1,2,3,4], target = 3

// Output: [1,2]
// Explanation:
// The sum of 1 and 2 is 3. Since we are assuming a 1-indexed array, index1 = 1, index2 = 2. We return [1, 2].

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

// Input: numbers = [5,25,75], target = 100
// Output: [2, 3]
// Explanation: The sum of 25 and 75 is 100. Therefore index1 = 2, index2 = 3. We return [2, 3].

// 2 <= numbers.length <= 1000
// -1000 <= numbers[i] <= 1000
// -1000 <= target <= 1000
package main

import (
	"fmt"
	"reflect"
)

func twoSumII(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}

		if numbers[left]+numbers[right] < target {
			left++
			continue
		}

		break
	}
	return []int{left + 1, right + 1}
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{1, 2, 3, 4}, 3, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{5, 25, 75}, 100, []int{2, 3}},
		{[]int{1, 2, 3, 4, 5, 7, 10, 11}, 9, []int{2, 6}},
	}

	for _, testCase := range testCases {
		res := twoSumII(testCase.nums, testCase.target)

		if !reflect.DeepEqual(res, testCase.result) {
			fmt.Printf("Test failed: %v and %v, Expected: %v but Got: %v\n", testCase.nums, testCase.target, testCase.result, res)
		}else {
			fmt.Printf("Test passed: %v and %v\n", testCase.nums, testCase.target)
		}
	}
}
