package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(h *ListNode) {
	for h != nil {
		fmt.Printf("%d -> ", h.Val)
		h = h.Next
	}
	fmt.Println("")
}

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	cur := head
	prev := newHead

	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		temp := cur.Next.Next
		cur.Next.Next = cur
		cur.Next = temp
		prev = cur
		cur = temp
	}

	return newHead.Next
}

func main() {
	A := new(ListNode)
	B := new(ListNode)
	C := new(ListNode)
	D := new(ListNode)
	A.Val = 1
	A.Next = B
	B.Val = 2
	B.Next = C
	C.Val = 3
	C.Next = D
	D.Val = 4
	h := swapPairs(A)
	printListNode(h)
}
