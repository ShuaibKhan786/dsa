// Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.
// Input: nums = [1, 2, 3, 3]

// Output: true

//hashmap problem

package main

import "fmt"


func hasDuplicate(nums []int) bool {
	hashTable := make(map[int]struct{}, 0)

	for _, num := range nums {
		if _, ok := hashTable[num]; ok {
			return true
		}
		hashTable[num] = struct{}{}
	}

	return false
}

func main() {
	testCases := []struct{
		nums   []int
		result bool
	}{
		{nums: []int{1, 2, 3, 3}, result: true},
		{nums: []int{1, 2, 3, 4}, result: false},
	}
	
	for _, testCase := range testCases {
		if state := hasDuplicate(testCase.nums); state != testCase.result {
			fmt.Printf("Test Case: %v failed, Expected: %v, Got: %v\n", testCase.nums, testCase.result, state)
		}else{
			fmt.Printf("Test Case: %v passed\n", testCase.nums)
		}
	}
}