package main

import (
	"log" 
	"container/heap"
)

type heapInt []int

func(h heapInt) Len() int {
	return len(h)
}

func(h heapInt) Less(i, j int) bool {
	return h[i] < h[j]
}

func(h heapInt) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func(h *heapInt) Push(el interface{}) {
	*h = append(*h, el.(int))
}

func(h *heapInt) Pop() interface{} {
	before := *h 
	lastElement := before[ before.Len() - 1 ]
	*h = before[0 : before.Len() - 1]
	return lastElement
}

func main() {
	intHeap := &heapInt{1,2,3}
	heap.Init(intHeap)
	heap.Push(intHeap, 4)
	heap.Push(intHeap, 5)
	heap.Push(intHeap, 6)

	heap.Pop(intHeap)
	heap.Pop(intHeap)
	heap.Pop(intHeap)
	heap.Pop(intHeap)
	
	log.Println(intHeap)
}