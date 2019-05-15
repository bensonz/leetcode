package main

import (
	"fmt"
)

/*
solution idea
use a linked list


OK so this solution doesn't really work. And I'll just come back later when things become clear
*/

type vlink struct {
	start int
	end   int
	next  *vlink
}

func listToLink(a []int) vlink {
	return vlink{
		start: a[0],
		end:   a[1],
		next:  nil,
	}
}

func addAndDrop(head *vlink, e *vlink) {
	shouldAdd := (e.start > head.end || e.end > head.end)
	switch {
	case head.next == nil && shouldAdd:
		// no next, simply put it behind
		fmt.Println("Adding : ", *e)
		head.next = e
	case e.start < head.start:
		// e covers area head does not cover, replace head
		e.next = head.next
		*head = *e
	case e.start == head.start && e.end > head.end:
		// also replace
		e.next = head.next
		*head = *e
	case shouldAdd:
		// e also covers area head cannot cover, add it to next
		addAndDrop(head.next, e)
	default:
		// drop e
		fmt.Println("Dropping : ", *e)
		return
	}
}

func videoStitching(clips [][]int, T int) int {
	var max = clips[0][1]
	var min = clips[0][0]
	// add the first element
	var w vlink
	for _, value := range clips {
		if max < value[1] {
			max = value[1]
		}
		if min > value[0] {
			min = value[0]
		}
		e := listToLink(value)
		addAndDrop(&w, &e)
	}
	if min != 0 || max < T {
		return -1
	}
	k := -1
	for p := &w; p != nil; p = p.next {
		fmt.Printf("(%d,%d) -> ", p.start, p.end)
		k++
		if p.end >= T {
			break
		}
	}
	fmt.Printf("\nk = %d\n", k)
	return k
}

func v(c [][]int, T int) int {
	return videoStitching(c, T)
}

func main() {
	c := [][]int{{0, 1}, {0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	fmt.Println(v(c, 10) == 3)
	fmt.Println(v(c, 5) == 2)
	fmt.Println(v(c, 11) == -1)
	b := [][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}
	fmt.Println(v(b, 9) == 3)
}
