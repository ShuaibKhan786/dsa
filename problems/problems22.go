// You are given an integer array nums consisting of n elements, and an integer k.

// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000

// Constraints:

// n == nums.length
// 1 <= k <= n <= 105
// -104 <= nums[i] <= 104

package main

import "fmt"

func findMaxAverage(nums []int, ws int) float64 {
	sum := 0

	for i := 0; i < ws; i++ {
		sum += nums[i]
	}

	maxSum := sum

	left := 1
	right := ws

	for right < len(nums) {
		newSum := (sum - nums[left-1]) + nums[right]

		if newSum > maxSum {
			maxSum = newSum
		}

		sum = newSum

		left++
		right++
	}

	return float64(maxSum) / float64(ws)
}

func main() {
	testCases := []struct {
		nums   []int
		ws     int
		output float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00},
	}

	for _, testCase := range testCases {
		result := findMaxAverage(testCase.nums, testCase.ws)

		if result != testCase.output {
			fmt.Printf("Test Failed: %v and %v, Expected: %v and Got: %v\n", testCase.nums, testCase.ws, testCase.output, result)
		} else {
			fmt.Printf("Test Passed: %v and %v\n", testCase.nums, testCase.ws)
		}
	}
}
