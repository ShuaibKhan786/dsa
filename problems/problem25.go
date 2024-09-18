// Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".

// You may assume that the correct output is always unique.

// Example1:
// input: s = "OUZODYXAZV", t = "XYZ"
// Output: "YXAZ"
// Explanation: "YXAZ" is the shortest substring that includes "X", "Y", and "Z" from string t.

// Example 2:
// input: s = "xyz", t = "xyz"
// Output: "xyz"

// Example 3:
// input: s = "x", t = "xy"
// Output: ""

// Constraints:

// 1 <= s.length <= 1000
// 1 <= t.length <= 1000
// s and t consist of uppercase and lowercase English letters.

package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	needM := make(map[byte]int)
	for i := range t {
		needM[t[i]]++
	}

	windowM := make(map[byte]int)

	left, right := 0, 0
	have, need := 0, len(needM)
	minLen := len(s) + 1
	minStrIndex := [2]int{-1, -1}

	for right < len(s) {
		charRight := s[right]

		if _, ok := needM[charRight]; ok {
			windowM[charRight]++
			if windowM[charRight] == needM[charRight] {
				have++
			}
		}

		for have == need {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
				minStrIndex[0] = left
				minStrIndex[1] = right
			}

			charLeft := s[left]

			if _, ok := needM[charLeft]; ok {
				windowM[charLeft]--
				if windowM[charLeft] < needM[charLeft] {
					have--
				}
			}

			left++
		}

		right++
	}

	if minStrIndex[0] == -1 {
		return ""
	}

	return s[minStrIndex[0] : minStrIndex[1]+1]
}

func main() {
	testCases := []struct {
		s      string
		t      string
		output string
	}{
		{"OUZODYXAZV", "XYZ", "YXAZ"},
		{"xyz", "xyz", "xyz"},
		{"ADOBECODEBANC", "ABC", "BANC"},
	}

	for _, testCase := range testCases {
		result := minWindow(testCase.s, testCase.t)

		if result != testCase.output {
			fmt.Printf("Test Failed: %v and %v, Expected: %v and Got: %v\n", testCase.s, testCase.t, testCase.output, result)
		} else {
			fmt.Printf("Test Passed: %v and %v\n", testCase.s, testCase.t)
		}
	}
}
