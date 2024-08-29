// You are given an array of integers temperatures where temperatures[i] represents the daily temperatures on the ith day.

// Return an array result where
// 	result[i] is the number of days after the ith day before a warmer temperature appears on a future day.
// 	If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

// Example 1:
// input: temperatures = [30,38,30,36,35,40,28]
// Output: [1,4,1,2,1,0,0]

// input: temperatures = [22,21,20]
// Output: [0,0,0]

// Constraints:

// 1 <= temperatures.length <= 1000.
// 1 <= temperatures[i] <= 100

package main

import (
	"fmt"
	"reflect"
)

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	s := make(Stack[int], 0)

	result[len(temperatures)-1] = 0
	s.Push(len(temperatures) - 1)
	for i := len(temperatures) - 2; i >= 0; i-- {
		for !s.IsEmpty() && temperatures[i] > temperatures[s.Peek()] {
			s.Pop()
		}

		if s.IsEmpty() {
			result[i] = 0
		} else {
			result[i] = s.Peek() - i
		}

		s.Push(i)
	}

	return result
}

func main() {
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{30, 38, 30, 36, 35, 40, 28}, []int{1, 4, 1, 2, 1, 0, 0}},
		{[]int{22, 21, 20}, []int{0, 0, 0}},
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, testCase := range testCases {
		result := dailyTemperatures(testCase.input)
		if !reflect.DeepEqual(testCase.output, result) {
			fmt.Printf("Test Failed: %v, Expected: %v but Got: %v\n", testCase.input, testCase.output, result)
		} else {
			fmt.Printf("Test Passed: %v\n", testCase.input)
		}
	}
}
