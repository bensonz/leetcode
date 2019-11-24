package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	if l == nil {
		fmt.Printf("Empty\n")
		return
	}
	for l != nil {
		fmt.Printf("%v -> ", l.Val)
		l = l.Next
	}
	fmt.Printf("nil\n")
}

func mergeList(l1, l2 *ListNode) *ListNode {
	var cur *ListNode
	var intersect bool
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l2
	}
	cur = &ListNode{}
	dummy := cur
	for l1 != nil && l2 != nil {
		if l1 == l2 {
			// same node
			fmt.Println("Same node")
			intersect = true
			break
		}
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if intersect {
		// the rest is the same
		for l1 != nil {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}
		return dummy.Next
	}
	for l1 != nil {
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	fmt.Println("End")
	return dummy.Next
}

func test1() {
	var l1, l2 *ListNode
	fmt.Println("Test1:")
	t1 := mergeList(l1, l2)
	printList(t1)
}

func test2() {
	var l1, l2 *ListNode
	l2 = &ListNode{1, nil}
	fmt.Println("Test2:")
	printList(l1)
	printList(l2)
	t2 := mergeList(l1, l2)
	printList(t2)
}

func test3() {
	var l1, l2 *ListNode
	l2 = &ListNode{1, nil}
	l1 = &ListNode{2, nil}
	fmt.Println("Test3:")
	printList(l1)
	printList(l2)
	t2 := mergeList(l1, l2)
	printList(t2)
}

func test4() {
	var l1, l2 *ListNode
	l2 = &ListNode{1, nil}
	l1 = &ListNode{2, nil}
	l1.Next = &ListNode{3, nil}
	l1.Next.Next = &ListNode{4, nil}
	fmt.Println("Test4:")
	printList(l1)
	printList(l2)
	t4 := mergeList(l1, l2)
	printList(t4)
}

func test5() {
	var l1, l2 *ListNode
	l1 = &ListNode{2, nil}
	l1.Next = &ListNode{3, nil}
	l1.Next.Next = &ListNode{4, nil}
	l2 = l1
	fmt.Println("Test5:")
	printList(l1)
	printList(l2)
	t4 := mergeList(l1, l2)
	printList(t4)
}

func test6() {
	var l1, l2 *ListNode
	l1 = &ListNode{2, nil}
	l1.Next = &ListNode{3, nil}
	l1.Next.Next = &ListNode{4, nil}
	neq := &ListNode{6, nil}
	neq.Next = &ListNode{7, nil}
	l1.Next.Next.Next = neq
	l2 = &ListNode{1, nil}
	l2.Next = neq
	fmt.Println("Test6:")
	printList(l1)
	printList(l2)
	t4 := mergeList(l1, l2)
	printList(t4)
}

func test() {
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
}

func main() {
	test()
}
