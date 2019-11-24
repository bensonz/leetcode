package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push ...
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop ...
func (h *IntHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func findKthSmall(nums []int, k int) int {
	minHeap := &IntHeap{}
	for i, v := range nums {
		if i < k {
			// while not enough elements, keep pushing
			minHeap.push(v)
		} else {
			if v < (*minHeap)[0] {
				// should push
				minHeap.push(v)
			}
		}
	}
}

func main() {
	fmt.Println()
}
