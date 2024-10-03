package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLL struct {
	Head *Node
	Tail *Node
	Len int
}

func NewSLL() *SinglyLL {
	return &SinglyLL{
		Head: nil,
		Tail: nil,
		Len: 0,
	}
}

func (ll *SinglyLL) Insert(item int) {
	n := new(Node)
	n.Data = item
	n.Next = nil

	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		ll.Tail = n
	}

	ll.Len++
}

func (ll *SinglyLL) Delete(item int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Data == item {
		ll.Head = ll.Head.Next
		ll.Len--
		return
	}

	prev := ll.Head
	for temp := prev.Next; temp != nil; temp = temp.Next {
		if temp.Data == item {
			prev.Next = temp.Next
			ll.Len--
			return
		}
		prev = temp
	}
}

func (ll *SinglyLL) Search(item int) bool {
	for temp := ll.Head; temp != nil; temp = temp.Next {
		if temp.Data == item {
			return true
		}
	}

	return false
}

func (ll *SinglyLL) Traverse() {
	for temp := ll.Head; temp != nil; temp = temp.Next {
		fmt.Printf("%d<-", temp.Data)
	}
	fmt.Println()
}

func (ll *SinglyLL) Length() int {
	return ll.Len
}

// this method is similar to reverseLinkList function where 
// head node is given and do the reverse logic and then return the head node
func (ll *SinglyLL) Reverse() {
	ll.Tail = ll.Head
	var prev *Node = nil 

	for ll.Head != nil {
		cachedNext := ll.Head.Next

		ll.Head.Next = prev
		prev = ll.Head

		ll.Head = cachedNext
	}

	ll.Head = prev
}

//here two ASC sorted list1 and list2 singlylist in given
//and I need to merge it
func mergeTwoSortedLL(list1, list2 *Node) *Node {
	var head *Node = nil
	var tail *Node = nil

	for list1 != nil && list2 != nil {
		if list1.Data < list2.Data {
			if head == nil {
				head = list1
				tail = head
			} else {
				tail.Next = list1
				tail = tail.Next
			}
			list1 = list1.Next
		} else {
			if head == nil {
				head = list2
				tail = head
			} else {
				tail.Next = list2
				tail = tail.Next
			}
			list2 = list2.Next
		}
	}

	for list1 != nil {
		if head == nil {
			return list1
		}
		tail.Next = list1
		list1 = list1.Next
	}

	for list2 != nil {
		if head == nil {
			return list2
		}
		tail.Next = list2
		list2 = list2.Next
	}

	return head
}

func main() {

	//reversing the linklist solution
	sll := NewSLL()
	sll.Insert(20)
	sll.Delete(20)
	sll.Insert(20)
	sll.Insert(30)
	sll.Insert(10)
	sll.Insert(10)

	sll.Traverse()

	sll.Reverse()

	sll.Traverse()

	//merging two sorted linklist solution
	sll1 := NewSLL()
	sll1.Insert(1)
	sll1.Insert(2)
	sll1.Insert(3)
	sll1.Insert(5)

	sll2 := NewSLL()
	sll2.Insert(4)
	sll2.Insert(6)
	sll2.Insert(7)

	mergedLL := *&SinglyLL{
		Head: mergeTwoSortedLL(sll1.Head, sll2.Head),
		Tail: nil,
		Len: 0,
	}

	mergedLL.Traverse()
}