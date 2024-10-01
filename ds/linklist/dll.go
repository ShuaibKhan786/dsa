package main

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type Dll struct {
	Head *Node
	Tail *Node
	Len  int
}

func NewDll() *Dll {
	return &Dll{
		Head: nil,
		Tail: nil,
		Len:  0,
	}
}

func (dll *Dll) InsertBegin(data int) {
	node := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}

	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
		dll.Len++
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node

		dll.Head = node
	}
}

func (dll *Dll) InsertEnd(data int) {
	node := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}

	if dll.Tail == nil {
		dll.Head = node
		dll.Tail = node
		dll.Len++
	} else {
		node.Prev = dll.Tail
		dll.Tail.Next = node

		dll.Tail = node
	}
}

func (dll *Dll) Insert(pos, data int) {

}

func (dll *Dll) DeleteHead() bool {
	if dll.Head != nil {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
		return true
	}
	return false
}

func (dll *Dll) DeleteTail() bool {
	if dll.Tail != nil {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	}
	return false
}

func (dll *Dll) Delete(data int) bool {
	for temp := dll.Head; temp != nil; temp = temp.Next {
		if temp.Data == data {
			prev := temp.Prev
			next := temp.Next

			switch {
			case prev == nil: 
				dll.Head = next
			case next == nil:
				dll.Tail = prev
			default:
				prev.Next = next
				next.Prev = prev
			}

			return true
		}
	}
	return false
}

func (dll *Dll) TraverseForward() {
	for temp := dll.Head; temp != nil; temp = temp.Next {
		fmt.Printf("%d->", temp.Data)
	}
	fmt.Println()
}

func (dll *Dll) TraverseReverse() {
	for temp := dll.Tail; temp != nil; temp = temp.Prev {
		fmt.Printf("<-%d", temp.Data)
	}
	fmt.Println()
}

func (dll *Dll) Search(data int) bool {
	for temp := dll.Tail; temp != nil; temp = temp.Prev {
		if temp.Data == data {
			return true
		}
	}
	return false
}

func (dll *Dll) Length() int {
	return dll.Len
}

func main() {
	dll := NewDll()

	dll.InsertBegin(2)
	dll.InsertBegin(3)
	dll.InsertEnd(4)
	dll.InsertEnd(5)

	dll.TraverseForward()
	dll.TraverseReverse()

	dll.InsertBegin(6)
	dll.InsertEnd(7)

	dll.TraverseForward()
	dll.TraverseReverse()

	dll.DeleteHead()
	dll.DeleteTail()

	dll.TraverseForward()
	dll.TraverseReverse()

	fmt.Println(dll.Search(2))

	dll.Delete(2)

	fmt.Println(dll.Search(2))

	dll.TraverseForward()
	dll.TraverseReverse()
}
