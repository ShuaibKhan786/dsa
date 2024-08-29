// You are given a string s consisting of the following
// characters: '(', ')', '{', '}', '[' and ']'.

// The input string s is valid if and only if:

// Every open bracket is closed by the same type of close bracket.
// Open brackets are closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
// Return true if s is a valid string, and false otherwise.

// Example 1:
// input: s = "[]"Output: true

// Example 2:
// input: s = "([{}])"Output: true

// Example 3:
// input: s = "[(])"Output: false

// Constraints:

// 1 <= s.length <= 1000

package main

import "fmt"

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	top := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]

	return top
}



func validParenthesis(str string) bool {
	s := make(Stack[rune], 0)

	hmap := map[rune]rune {
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range str {
		if match, ok := hmap[r]; ok {
			if s.Pop() != match {
				return false
			}
		}else {
			s.Push(r)
		}
	}

	return len(s) == 0
}

func main() {
	testCases := []struct {
		str    string
		result bool
	}{
		{"([{}])", true},
		{"[]", true},
		{"[(])", false},
	}

	for _, testCase := range testCases {
		res := validParenthesis(testCase.str)

		if res != testCase.result {
			fmt.Printf("Test failed: %s, Expected: %v but Got: %v\n", testCase.str, testCase.result, res)
		} else {
			fmt.Printf("Test passed: %s\n", testCase.str)
		}
	}
}
