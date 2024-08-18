// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
// Anagram : An anagram is a string that contains the exact same characters as another string,
// but the order of the characters can be different.

// Input: s = "racecar", t = "carrace"

// Output: true

// Input: s = "jar", t = "jam"

// Output: false

package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[rune]int, 0), make(map[rune]int, 0)

	for _, r := range s {
		sMap[r] = sMap[r] + 1
	}

	for _, r := range t {
		tMap[r] = tMap[r] + 1
	}

	for _, r := range s {
		if sMap[r] != tMap[r] {
			return false
		}
	}

	return true
}

func main() {
	testCases := []struct {
		s      string
		t      string
		result bool
	} {
		{ "foo", "ofo", true},
		{ "racecar", "carrace", true},
		{ "jar", "jam", false},
		{ "anagram", "nagaram", true},
	}

	for _, testCase := range testCases {
		if res := isAnagram(testCase.s, testCase.t); res != testCase.result {
			fmt.Printf("TestCase failed s=%s and t=%s, Expected: %v, Got: %v\n", testCase.s, testCase.t, testCase.result, res)
		}else {
			fmt.Printf("TestCase Passed s=%s and t=%s\n", testCase.s, testCase.t)
		}
	}
}
