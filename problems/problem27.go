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

//TC: O(n)
//SC: O(1)
func removeNthFromEndBF(sll *SLL, n int) {
	length := 0

	//step1: vists all node once to find the length of a linklist
	for temp := sll.Head; temp != nil; temp = temp.Next {
		length++
	}

	//edge case no point of removing the n node which is greater than the length
	if length < n {
		return
	}

	//it means the node we need to remove is head node
	if length == n {
		sll.Head = sll.Head.Next
		return
	}

	//step2: travserse upto the node just before the node need to be removed
	temp := sll.Head
	for i := 1; i < length - n; i++ {
		temp = temp.Next
	}

	//removed the node
	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}
}

func removeNthFromEnd(sll *SLL, n int) {
	dummy := &Node{ Next: sll.Head}
	
	first := dummy
	second := dummy

	//first must be n head from first
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	//second will point to the node just before nth node we want to removed
	second.Next = second.Next.Next

	sll.Head = dummy.Next
}

//TC: O(n) + O(m) i,e O(n+m)
//SC: O(n)
//you can combined in one loop but
//for clarity I do a separate the edge case logic
func sumOfTwoNumber(sll1, sll2 *SLL) *SLL {
	l1 := sll1.Head
	l2 := sll2.Head

	var newHead *Node = nil
	var newTail *Node = nil
	carry := 0

	for l1 != nil && l2 != nil {
		sum := l1.Data + l2.Data + carry
		remainder := sum % 10
		carry = sum / 10

		newNode := &Node{
			Data: remainder,
			Next: nil,
		} 

		if newHead == nil {
			newHead = newNode
			newTail = newHead
		} else {
			newTail.Next = newNode
			newTail = newNode
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Data + carry
		remainder := sum % 10
		carry = sum / 10

		newNode := &Node{
			Data: remainder,
			Next: nil,
		}

		newTail.Next = newNode
		newTail = newNode

		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Data + carry
		remainder := sum % 10
		carry = sum / 10

		newNode := &Node{
			Data: remainder,
			Next: nil,
		}

		newTail.Next = newNode
		newTail = newNode

		l2 = l2.Next
	}

	for carry != 0 {
		remainder := carry % 10
		carry = carry / 10

		newNode := &Node{
			Data: remainder,
			Next: nil,
		}

		newTail.Next = newNode
		newTail = newNode
	}

	return &SLL{
		Head: newHead,
	}
}

func hasCycle(head *Node) bool {
	tortoise := head
	rabbit := head

	for tortoise != nil && rabbit != nil && rabbit.Next != nil {
		tortoise = tortoise.Next
		rabbit = rabbit.Next.Next

		if tortoise == rabbit {
			return true
		}
	}

	return false
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

	removeNthFromEnd(mergedSll, 5)
	mergedSll.Traverse()

	fmt.Println("*****Sum of two number*****")
	l1 := NewSLL()
	l1.Insert(9)
	l1.Insert(5)
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(4)
	l1.Traverse()

	l2 := NewSLL()
	l2.Insert(7)
	l2.Insert(9)
	l2.Insert(1)
	l2.Traverse()

	sum := sumOfTwoNumber(l1, l2)
	sum.Traverse()

	fmt.Println("*****Linked List cycle*****")
	n1 := &Node{ Data: 1, Next: nil}
	n2 := &Node{ Data: 2, Next: nil}
	n3 := &Node{ Data: 3, Next: nil}
	n4 := &Node{ Data: 4, Next: nil}

	head := n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	fmt.Println(hasCycle(head))
}
