package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SLL struct {
	Head *Node
}

func NewSLL() *SLL {
	return &SLL{
		Head: nil,
	}
}

func (sll *SLL) Insert(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if sll.Head == nil {
		sll.Head = newNode
	} else {
		newNode.Next = sll.Head
		sll.Head = newNode
	}
}

func (sll *SLL) Traverse() {
	for temp := sll.Head; temp != nil; temp = temp.Next {
		fmt.Printf("%d <- ", temp.Data)
	}
	fmt.Println()
}

// reverse a singly linklist
// TC = O(n), SC = O(1)
func reverseSLL(sll *SLL) {
	var prev *Node = nil

	for sll.Head != nil {
		cachedNext := sll.Head.Next

		sll.Head.Next = prev
		prev = sll.Head

		sll.Head = cachedNext
	}

	sll.Head = prev
}

//merging two sorted sll ASC
// TC = O(n+m), SC = O(1)
func mergeSLL(firstSLL *SLL, secondSLL *SLL) *SLL {
	var mergedHead *Node = nil 

	mergedTail := mergedHead
	firstHead := firstSLL.Head
	secondHead := secondSLL.Head

	for firstHead != nil && secondHead != nil {
		if firstHead.Data < secondHead.Data {
			if mergedHead == nil {
				mergedHead = firstHead
				mergedTail = mergedHead
			} else {
				mergedTail.Next = firstHead
				mergedTail = firstHead
			}
			firstHead = firstHead.Next
		} else {
			if mergedHead == nil {
				mergedHead = secondHead
				mergedTail = mergedHead
			} else {
				mergedTail.Next = secondHead
				mergedTail = secondHead
			}
			secondHead = secondHead.Next
		}
	}

	for firstHead != nil {
		mergedTail.Next = firstHead
		mergedTail = firstHead
		firstHead = firstHead.Next
	}

	for secondHead != nil {
		mergedTail.Next = secondHead
		mergedTail = secondHead
		secondHead = secondHead.Next
	}

	return &SLL{
		Head: mergedHead,
	}
}

//TC: 
//	O(n) for finding midpoint
//	O(n/2) i,e O(n) for reversing second half
//  O(n) for merging first and seconf half
//  thus it gets to O(n) + O(n) + O(n) = 3O(n) i,e O(n) we discard coefficient in bigO notation
//SC: O(1)
func reorder(sll *SLL) {
	//step1: find the midpoint
	slow := sll.Head
	fast := sll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//step2: reversed the second half
	var reverseHead *Node = nil
	for slow != nil {
		cachedNext := slow.Next

		slow.Next = reverseHead
		reverseHead = slow
		
		slow = cachedNext
	}

	//step3: merged them alternately
	first := sll.Head
	second := reverseHead

	for first != nil && second != nil {
		cachedFirst := first.Next
		cachedSecond := second.Next
	
		first.Next = second
		second.Next = cachedFirst
		
		first = cachedFirst
		second = cachedSecond
	}

	if first != nil { //odd condition
		first.Next = nil
	}
}

func main() {
	sll1 := NewSLL()
	sll1.Insert(1)
	sll1.Insert(3)
	sll1.Insert(7)
	sll1.Insert(9)
	sll1.Insert(10)
	sll1.Traverse()

	reverseSLL(sll1)
	sll1.Traverse()

	sll2 := NewSLL()
	sll2.Insert(1)
	sll2.Insert(2)
	sll2.Insert(3)
	sll2.Insert(4)
	sll2.Insert(5)
	sll2.Insert(6)
	sll2.Insert(8)
	sll2.Insert(11)
	sll2.Traverse()

	reverseSLL(sll2)
	sll2.Traverse()

	mergedSll := mergeSLL(sll1, sll2)
	mergedSll.Traverse()

	reorder(mergedSll)
	mergedSll.Traverse()
}
