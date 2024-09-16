// You are given an integer array prices where prices[i] is the price of NeetCoin on the ith day.

// You may choose a single day to buy one NeetCoin and choose a different day in the future to sell it.

// Return the maximum profit you can achieve. You may choose to not make any transactions, in which case the profit would be 0.

// Example 1:
// input: prices = [10,1,5,6,7,1]
// Output: 6
// Explanation: Buy prices[1] and sell prices[4], profit = 7 - 1 = 6.

// Example 2:
// input: prices = [10,8,7,5,2]
// Output: 0

// Explanation: No profitable transactions can be made, thus the max profit is 0.

// Constraints:

// 1 <= prices.length <= 100
// 0 <= prices[i] <= 100

package main

import "fmt"

func maxProfit(prices []int) int {
	buy := 0
	sell := 1
	maxP := 0

	for sell < len(prices) {
		// making sure we get profit when sell the stock
		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			maxP = max(maxP, profit)
		} else {
			buy = sell
		}

		sell++
	}

	return maxP
}

func main() {
	testCases := []struct {
		prices []int
		output int
	} {
		{ []int{10,1,5,6,7,1} , 6},
		{ []int{10,8,7,5,2} , 0},
		{ []int{7, 1, 5, 3, 6, 4} , 5},
	}

	for _, testCase := range testCases {
		result := maxProfit(testCase.prices)

		if result != testCase.output {
			fmt.Printf("Test Failed: %v, Expected: %v but Got: %v\n", testCase.prices, testCase.output, result)
		} else {
			fmt.Printf("Test Passed: %v\n", testCase.prices)
		}
	}
}
