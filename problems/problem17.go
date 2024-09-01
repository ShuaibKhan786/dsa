// You are given an m x n 2-D integer array matrix and an integer target.

// Each row in matrix is sorted in non-decreasing order.
// The first integer of every row is greater than the last integer of the previous row.
// Return true if target exists within matrix or false otherwise.

// Can you write a solution that runs in O(log(m * n)) time?

// Example 1:
// input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]],
// target = 10
// Output: true

// Example 2:
// input: matrix = [[1,2,4,8],[10,11,12,13],[14,20,30,40]],
// target = 15
// Output: false

package main

import "fmt"

func potentialRow(matrix [][]int, target int) (int, bool) {
	top := 0
	bottom := len(matrix) - 1

	for top <= bottom {
		mid := (top + bottom) / 2

		switch {
		case matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target:
			if matrix[mid][0] == target {
				return mid, true
			}
			return mid, false
		case matrix[mid][0] > target:
			bottom = mid - 1
		default:
			top = mid + 1
		}
	}

	return -1, false
}

func searchMatrix(matrix [][]int, target int) bool {
	pRow, ok := potentialRow(matrix, target)
	if ok {
		return true
	}

	if pRow == -1 {
		return false
	}

	left := 0
	right := len(matrix[pRow]) - 1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case matrix[pRow][mid] > target:
			right = mid - 1
		case matrix[pRow][mid] < target:
			left = mid + 1
		default:
			return true
		}
	}

	return false
}

func main() {
	testCases := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{[][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 15, false},
		{[][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {59, 60, 62, 66}, {79, 80, 83, 86}, {90, 91, 93, 97}}, 93, true},
	}

	for _, testCase := range testCases {
		result := searchMatrix(testCase.matrix, testCase.target)

		if result != testCase.output {
			fmt.Printf("Test failed: %v and %v , Expected: %v but Got: %v\n", testCase.matrix, testCase.target, testCase.output, result)
		} else {
			fmt.Printf("Test passed: %v and %v\n", testCase.matrix, testCase.target)
		}
	}
}
