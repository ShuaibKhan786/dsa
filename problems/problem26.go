// You are given an array of integers nums and an integer k. There is a sliding window of size k that starts at the left edge of the array.
// The window slides one position to the right until it reaches the right edge of the array.

// Return a list that contains the maximum element in the window at each step.

// input: nums = [1,2,1,0,4,2,6], k = 3
// Output: [2,2,4,4,6]
// Explanation: Window position
// Max---------------           -----
// [1  2  1] 0  4  2  6        2
// 1 [2  1  0] 4  2  6         2
// 1  2 [1  0  4] 2  6         4
// 1  2  1 [0  4  2] 6         4
// 1  2  1  0 [4  2  6]        6

package main

import "fmt"

//Brut force solution

// SC = O(n^2)
// TC = O(n)
func maxSlidingWindowBF(nums []int, ws int) []int  {
	maxArr := make([]int, 0)

	left := 0
	right := ws - 1

	for right < len(nums){
		currentMax := nums[left]

		tempIndex := left + 1 

		for tempIndex <= right {
			currentMax = max(currentMax, nums[tempIndex]) 
			tempIndex++
		}

		maxArr = append(maxArr, currentMax)

		left++
		right++
	}

	return maxArr
}

type MonotonicQueue struct {
	Data []int
}

func (q *MonotonicQueue) Enqueue(nums []int, index int) {
	for  len(q.Data) > 0 && nums[index] > nums[q.Data[len(q.Data) - 1]] {
		q.Data = q.Data[:len(q.Data) - 1]
	}

	q.Data = append(q.Data, index)
}

func (q *MonotonicQueue) Dequeue(index int) {
	if len(q.Data) > 0 && q.Data[0] < index {
		q.Data = q.Data[1:]
	}
}

func (q *MonotonicQueue) Max() int {
	if len(q.Data) > 0 {
		return q.Data[0]
	}
	return -1
}

func maxSlidingWindow(nums []int, ws int) []int  {
	result := []int{}

	mq := &MonotonicQueue{ Data: []int{}}
	for index  := range nums {
		mq.Enqueue(nums, index)

		mq.Dequeue(index - ws + 1)

		if index >= ws - 1 {
			result = append(result, nums[mq.Max()])
		}
	}

	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,2,1,0,4,2,6}, 3))
}