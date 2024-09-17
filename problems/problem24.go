// Given a string s, find the length of the longest substring without duplicate characters.

// A substring is a contiguous sequence of characters within a string.
// Example 1:
// input: s = "zxyzxyz"
// Output: 3

// Example 2:
// input: s = "xxxx"
// Output: 1

// Constraints:

// 0 <= s.length <= 1000
// s may consist of printable ASCII characters.

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]struct{})
	maxLength := 0

	left := 0
	for right, char := range s {
		for _, ok := set[char]; ok; _, ok = set[char] {
			delete(set, rune(s[left]))
			left++
		}

		set[char] = struct{}{}

		currentLength := right - left + 1

		maxLength = max(maxLength, currentLength)
	}

	return maxLength
}

func main() {
	testCases := []struct {
		input  string
		output int
	}{
		{"zxyzxyz", 3},
		{"xxxx", 1},
		{"pwwkew", 3},
		{"abcabcbb", 3},
	}

	for _, testCase := range testCases {
		result := lengthOfLongestSubstring(testCase.input)

		if result != testCase.output {
			fmt.Printf("Test Failed: %v, Expected: %v and Got: %v\n", testCase.input, testCase.output, result)
		} else {
			fmt.Printf("Test Passed: %v\n", testCase.input)
		}
	}
}
