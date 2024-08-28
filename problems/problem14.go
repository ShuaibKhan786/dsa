// You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.

// Return the integer that represents the evaluation of the expression.

// The operands may be integers or the results of other operations.
// The operators include '+', '-', '*', and '/'.
// Assume that division between integers always truncates toward zero.
// The answer and all the intermediate calculations can be represented in a 32-bit integer

// Example 1:
// input: tokens = ["1","2","+","3","*","4","-"]
// Output: 5
// Explanation: ((1 + 2) * 3) - 4 = 5
// 1 <= tokens.length <= 104
package main

import (
	"fmt"
	"log"
	"strconv"
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
	top := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]

	return top
}

func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	return (*s)[len(*s)-1]
}

func evalRPN(tokens []string) int {
	s := make(Stack[int], 0)

	for _, token := range tokens {
		switch token {
		case "+":
			operand1 := s.Pop()
			operand2 := s.Pop()

			s.Push(operand2 + operand1)
		case "-":
			operand1 := s.Pop()
			operand2 := s.Pop()

			s.Push(operand2 - operand1)
		case "*":
			operand1 := s.Pop()
			operand2 := s.Pop()

			s.Push(operand2 * operand1)
		case "/":
			operand1 := s.Pop()
			operand2 := s.Pop()

			s.Push(operand2 / operand1)
		default:
			tokenInt64, err := strconv.ParseInt(token, 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			s.Push(int(tokenInt64))
		}
	}

	return s.Pop()
}

func main() {
	testCases := []struct {
		tokens []string
		output int
	}{
		{[]string{"1", "2", "+", "3", "*", "4", "-"}, 5},
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, testCase := range testCases {
		result := evalRPN(testCase.tokens)

		if result != testCase.output {
			fmt.Printf("Test failed: %v, Expected: %v but Got: %v\n", testCase.tokens, testCase.output, result)
		}else {
			fmt.Printf("Test passed: %v\n", testCase.tokens)
		}
	}
}
