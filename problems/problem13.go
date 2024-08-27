// Design a stack class that supports the
// push, pop, top, and getMin operations.

// newXStack() initializes the stack object.
// xPush(int val) pushes the element val onto the stack.
// xPop() removes the element on the top of the stack.
// xTop() gets the top element of the stack.
// xGetMin() retrieves the minimum element in the stack.
// Each function should run in O(1) time.

// input: ["MinStack", "push", 1, "push", 2, "push", 0, "getMin", "pop", "top", "getMin"]
// Output: [null,null,null,null,0,null,2,1]
// Explanation:
// MinStack minStack = new MinStack();
// minStack.push(1);
// minStack.push(2);
// minStack.push(0);
// minStack.getMin();

// Constraints:

// -2^31 <= val <= 2^31 - 1.
// pop, top and getMin will always be called on non-empty stacks.

package main

import "fmt"

type XStack struct {
	NormStack []int
	MinStack  []int
}

func newXStack() *XStack {
	return &XStack{
		NormStack: make([]int, 0),
		MinStack:  make([]int, 0),
	}
}

func (s *XStack) xPush(value int) {
	if len(s.MinStack) == 0 {
		s.NormStack = append(s.NormStack, value)
		s.MinStack = append(s.MinStack, value)
	} else {
		s.NormStack = append(s.NormStack, value)
		if min := s.MinStack[len(s.MinStack)-1]; min < value {
			s.MinStack = append(s.MinStack, min)
		} else {
			s.MinStack = append(s.MinStack, value)
		}
	}
}

func (s *XStack) xPop() {
	if len(s.NormStack) == 0 {
		return
	}
	s.NormStack = s.NormStack[:len(s.NormStack)-1]
	s.MinStack = s.MinStack[:len(s.MinStack)-1]
}

func (s *XStack) xTop() int {
	return s.NormStack[len(s.NormStack)-1]
}

func (s *XStack) xGetMin() int {
	return s.MinStack[len(s.MinStack)-1]
}

func main() {
	xs := newXStack()
	xs.xPush(1)
	xs.xPush(2)
	xs.xPush(0)
	fmt.Println(xs.xGetMin()) // 0
	xs.xPop()
	fmt.Println(xs.xTop())    // 2
	fmt.Println(xs.xGetMin()) //1
}
