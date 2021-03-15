package main

import "fmt"

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
	start  Node
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
func (h *LinkedList) Push(v int) *LinkedList {
	if h.length == 0 {
		h.start = Node{value: v}
		h.end = &h.start
	} else {
		h.end.next = &Node{value: v}
		h.end = h.end.next
	}

	h.length++

	return h
}

/*
*	traverse the linked list using a callback
	@returns void
*/
func (h *LinkedList) Traverse(f func(*Node)) {
	var n *Node = &h.start

	for n != nil {
		f(n)
		n = n.next
	}
}

/*
*	removes the last element of list
 */
func (h *LinkedList) Pop() *LinkedList {

	return h
}

/*
*	main
 */
func main() {
	l := MakeList()

	l.Push(1).Push(2).Push(3).Push(100)

	l.Traverse(func(n *Node) {
		fmt.Println(n.value)
	})

	//fmt.Println(l.start.next.value)
}
