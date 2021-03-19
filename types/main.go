package main

import (
	"fmt"

	"github.com/tgbv/go-linkedlist/internal/listTypesSingly"
	ll "github.com/tgbv/go-linkedlist/pkg"
)

/*
*	base node
 */
type Node struct {
	value int
	next  *Node
}

/*
*	the list
 */
type LinkedList struct {
	length uint8
	start  *Node
	end    *Node
}

/*
*	creates list
 */
func MakeList() LinkedList {
	return LinkedList{length: 0}
}

/*
*	pushes a new element to nodes list
 */
func (h *LinkedList) Add(v int) *LinkedList {
	if h.length == 0 {
		h.start = &Node{value: v}
		h.end = h.start
	} else {
		h.end.next = &Node{value: v}
		h.end = h.end.next
	}

	h.length++

	return h
}

/*
*	lookup node at index
 */
func (h *LinkedList) Lookup(i uint8) *Node {
	//fmt.Println(h.length, i)
	// check for index boundaries
	if i < 0 || i > h.length-1 {
		panic("Index out of boundaries!")
	}

	// if index is 0 return head
	if i == 0 {
		return h.start
	} else

	// otherwise loop
	{
		var (
			j    uint8 = 1 // starts at 1 because we've already registered node = h.start aka one iteration
			node *Node = h.start
		)

		for j <= i {
			node = node.next
			j++
		}

		return node
	}
}

/*
*	deletes head
 */
func (h *LinkedList) DeleteHead() *LinkedList {
	// check length first
	// if it's zero, nothing must be done
	if h.length == 0 {
		return h
	}

	// otherwise make head point to its next
	h.start = h.start.next

	h.length--

	return h
}

/*
*	deletes node at index
	@return *LinkedList
*/
func (h *LinkedList) DeleteAtIndex(x uint8) *LinkedList {
	// check if index is in boundaries
	if x < 0 || x > h.length-1 {
		panic("Index out of boundaries !")
	}

	// if index is 0 just delete the head
	if x == 0 {
		return h.DeleteHead()
	} else

	// otherwise loop through nodes
	{
		var i uint8
		var prev *Node
		var curr *Node = h.start

		for i < h.length {
			if i == x {
				prev.next = curr.next
				break
			} else {
				prev = curr
				curr = prev.next
			}
			i++
		}

		h.length--

		return h
	}
}

/*
*	deletes tail
 */
func (h *LinkedList) DeleteTail() *LinkedList {
	// check length first
	// if it's 0 nothing must be done
	if h.length == 0 {
		return h
	}

	// otherwise delete at index of length-1
	return h.DeleteAtIndex(h.length - 1)
}

/*
*	traverse the linked list using a callback
	@returns void
*/
func (h *LinkedList) Traverse(f func(*Node)) {
	n := h.start

	for n != nil {
		f(n)
		n = n.next
	}
}

/*
*	main
 */
func main() {
	// l := MakeList()

	// l.Add(1).Add(2).Add(3).Add(100).DeleteHead().DeleteTail()

	// l.Traverse(func(n *Node) {
	// 	fmt.Println(n.value)
	// })

	//fmt.Println(l.start.next.value)

	l := ll.MakeSinglyListInt32Length32()

	l.Add(11).Add(78).Traverse(func(n *listTypesSingly.Int32Length32Node) {
		fmt.Println(n)
	})
}
