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

func main() {
	sll := NewSLL()

	fmt.Println(sll.Length())

	sll.Insert(20)
	sll.Delete(20)
	sll.Insert(20)
	sll.Insert(30)
	sll.Insert(10)
	sll.Insert(10)

	sll.Traverse()
	fmt.Println(sll.Len)

	sll.Delete(10)
	sll.Delete(10)

	sll.Traverse()
	fmt.Println(sll.Length())
}