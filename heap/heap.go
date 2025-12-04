package main

// https://replit.com/@eduardolfalcao/BinaryHeap?v=1#main.go

type PQ interface{
	Add(e int)
	Poll() (int, error)
	Remove(e int) error
	Init()
	doubleV()
	
}

type BinaryHeap struct {
	v               []int
	elementsInserted int
}

func (heap *BinaryHeap) indexOfParent(index int) int {
	return (index - 1) / 2
}

func (heap *BinaryHeap) indexOfChildren(index int) (int,int) {
	left_child := 2*index + 1
	right_child := 2*index + 2
	return left_child, right_child
}

func (heap *BinaryHeap) Add(value int) {
	heap.v[heap.elementsInserted] = value
	heap.elementsInserted++
	heap.maxHeapify(heap.elementsInserted-1)
}