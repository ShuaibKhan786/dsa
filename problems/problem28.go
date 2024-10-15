package main

import (
	"fmt"
)

//TC: O(n^2) SC: O(1)
func findDuplicateBF(nums []int) int {
	for outer := 0; outer < len(nums); outer++ {
		for inner := outer + 1; inner < len(nums); inner++ {
			if nums[outer] == nums[inner] {
				return nums[outer]
			}
		}
	}

	return 0
}

//TC: O(n) SC: O(n)
func findDuplicateHM(nums []int) int {
	//using hashset
	hset := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := hset[num]; ok {
			return num
		}
		hset[num] = struct{}{}
	}

	return 0
}

//TC: O(n) SC: O(1)
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
	  slow = nums[slow]
    fast = nums[nums[fast]]
  }

  slow = 0
  for slow != fast {
    slow = nums[slow]
    fast = nums[fast]
  }
  
  return slow
}

func main() {
	fmt.Println(findDuplicateBF([]int{1, 2, 3, 2, 2}))
	fmt.Println(findDuplicateHM([]int{1, 2, 3, 2, 2}))
	fmt.Println(findDuplicateBF([]int{1, 2, 3, 4, 4}))
	fmt.Println(findDuplicateHM([]int{1, 2, 3, 4, 4}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 4}))
	fmt.Println(findDuplicate([]int{1, 2, 3, 2, 2 }))
}
