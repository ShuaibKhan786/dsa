// Given an integer array nums and an integer k,
// return the k most frequent elements within the array.

// Input: nums = [1,2,2,3,3,3], k = 2

// Output: [2,3]

// Input: nums = [7,7], k = 1

// Output: [7]
package main

import (
	"fmt"
	"reflect"
)

func topKElements(nums []int, k int) []int {
	var highestFreq int

	hmap := make(map[int]int, 0)

	for _, num := range nums {
		hmap[num]++
		if hmap[num] > highestFreq {
			highestFreq = hmap[num]
		}
	}

	bucket := make([][]int, highestFreq+1)
	for i := range bucket {
		bucket[i] = make([]int, 0)
	}

	for elem, freq := range hmap {
		bucket[freq] = append(bucket[freq], elem)
	}

	result := make([]int, 0)
	for i := highestFreq; i >= 0 && k > 0; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
			k--
		}
	}

	return result
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 2, []int{3, 2}},
		{[]int{7, 7}, 1, []int{7}},
		{[]int{2, 2, 4, 4}, 2, []int{4, 2}},
		{[]int{1}, 1, []int{1}},
	}

	for _, testCase := range testCases {
		res := topKElements(testCase.nums, testCase.k)
		if !reflect.DeepEqual(res, testCase.result) {
			fmt.Printf("Test Case: %v failed, Expected: %v , Got: %v\n", testCase.nums, testCase.result, res)
		} else {
			fmt.Printf("Test Case: %v passed\n", testCase.nums)
		}
	}
}
