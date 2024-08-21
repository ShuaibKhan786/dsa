// Given a string s, return true if it is a palindrome, otherwise return false.

// A palindrome is a string that reads the same forward and backward.
// It is also case-insensitive and ignores all non-alphanumeric characters.

// Input: s = "Was it a car or a cat I saw?"

// Output: true

// Explanation:
// After considering only alphanumerical characters we have "wasitacaroracatisaw", which is a palindrome.

// Input: s = "tab a cat"

// Output: false

// Constraints:

// 1 <= s.length <= 1000
// s is made up of only printable ASCII characters.

package main

import (
	"fmt"
	"strings"
)

func isAlphaNumeric(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9'
}

func isPlaindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if !isAlphaNumeric(s[start]) {
			start++
			continue
		}

		if !isAlphaNumeric(s[end]) {
			end--
			continue
		}

		if strings.ToLower(string(s[start])) != strings.ToLower(string(s[end])) {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	testCases := []struct{
		s string
		result bool
	} {
		{ "Was it a car or a cat I saw?", true },
		{ "tab a cat", false},
		{ "A", true},
		{ "9", true},
	}

	for _, testCase := range testCases {
		res := isPlaindrome(testCase.s)
		if res != testCase.result {
			fmt.Printf("Test Failed %v, Expected: %v and Got: %v\n", testCase.s, testCase.result, res)
		} else {
			fmt.Printf("Test Passed %v\n", testCase.s)
		}
	}
}