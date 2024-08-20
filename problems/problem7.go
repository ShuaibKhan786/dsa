// Given an array of integers nums, return the length of the longest consecutive sequence of elements.

// A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element.

// You must write an algorithm that runs in O(n) time.

// Input: nums = [2,20,4,10,3,4,5]

// Output: 4

// Input: nums = [0,3,2,5,4,6,1,1]

// Output: 7

// Constraints:

// 0 <= nums.length <= 1000
// -10^9 <= nums[i] <= 10^9

package main

import "fmt"

func longestConsecutive(nums []int) int {
	hmap := make(map[int]bool)

	for _, num := range nums {
		hmap[num] = true
	}

	longestOfAll := 1

	for _, num := range nums {
		if !hmap[num-1] {
			currentLongest := 1
			currentNum := num + 1

			for hmap[currentNum] {
				currentLongest++
				currentNum++
			}

			if currentLongest > longestOfAll {
				longestOfAll = currentLongest
			}
		}
	}

	return longestOfAll
}

func main() {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{0, 3, 2, 5, 4, 6, 1, 1}, 7},
		{[]int{2,20,4,10,3,4,5}, 4},
		{[]int{0, -1, -2, 9, 1, 2}, 5},
	}

	for _, testCase := range testCases {
		res := longestConsecutive(testCase.nums)
		if res != testCase.result {
			fmt.Printf("Test Failed %v, Expected: %d and Got: %v\n", testCase.nums, testCase.result, res)
		}else {
			fmt.Printf("Test Passed %v\n", testCase.nums)
		}
	}
}
