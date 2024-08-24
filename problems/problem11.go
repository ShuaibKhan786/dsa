// You are given an integer array heights where heights[i] represents the height of the ith bar.

// You may choose any two bars to form a container. Return the maximum amount of water a container can store.

// Example 1:
// input: height = [1,7,2,5,4,7,3,6]Output: 36

// Example 2:
// input: height = [2,2,2]Output: 4

// Constraints:

// 2 <= height.length <= 1000
// 0 <= height[i] <= 1000

package main

import "fmt"

func MaxWaterContainer(height []int) int {
	maxArea := 0

	left := 0
	right := len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left) // area = height * width

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	testCases := []struct {
		height []int
		area   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 7, 2, 5, 4, 7, 3, 6}, 36},
		{[]int{2, 2, 2}, 4},
	}

	for _, testCase := range testCases {
		res := MaxWaterContainer(testCase.height)

		if res != testCase.area {
			fmt.Printf("Test failed: %v, Expected: %v but Got: %v\n", testCase.height, testCase.area, res)
		}else {
			fmt.Printf("Test passed: %v\n", testCase.height)
		}
	}
}
