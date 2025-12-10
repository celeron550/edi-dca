package main

import (
	"errors"
)

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

func (heap *BinaryHeap) bubbleDown(index int) {
	left, right := heap.indexOfChildren(index)
	smallest := index

	if left < heap.elementsInserted && heap.v[left] < heap.v[smallest] {
		smallest = left
	}
	if right < heap.elementsInserted && heap.v[right] < heap.v[smallest] {
		smallest = right
	}

	// se n trocou, para
	if smallest == index {
		return
	}

	heap.v[index], heap.v[smallest] = heap.v[smallest], heap.v[index]

	heap.bubbleDown(smallest)
}

func (heap *BinaryHeap) maxHeapify(index int) {
	// se o pai eh menor que o filho, troca
	if index >=0 && index < heap.elementsInserted && heap.v[heap.indexOfParent(index)] < heap.v[index]{
		heap.v[heap.indexOfParent(index)], heap.v[index] = heap.v[index], heap.v[heap.indexOfParent(index)]
		heap.maxHeapify(heap.indexOfParent(index))
	}
}

func (heap *BinaryHeap) Add(value int) {
	if heap.elementsInserted == len(heap.v) {
    	heap.v = append(heap.v, value)
	}
	heap.v[heap.elementsInserted] = value
	heap.elementsInserted++
	heap.maxHeapify(heap.elementsInserted-1)
}

func (heap *BinaryHeap) Poll() (int,error) {
	if heap.elementsInserted == 0 {
		return -1,errors.New("Heap vazia")
	}
	removed := heap.v[0]
    heap.v[0] = heap.v[heap.elementsInserted-1]
    heap.elementsInserted--
    heap.bubbleDown(0)	
	return removed, nil
	
}
func (heap *BinaryHeap) Remove(e int) error {
	if heap.elementsInserted == 0 {
        return errors.New("heap vazia")
    }

    // 1. encontrar o índice do elemento e
    index := -1
    for i := 0; i < heap.elementsInserted; i++ {
        if heap.v[i] == e {
            index = i
            break
        }
    }

    if index == -1 {
        return errors.New("elemento nao encontrado")
    }

    // 2. substituir pelo último elemento da heap
    heap.v[index] = heap.v[heap.elementsInserted-1]
    heap.elementsInserted--

    // 3. tentar subir (caso o novo valor seja maior que o pai)
    heap.maxHeapify(index)

    // 4. tentar descer (caso o novo valor seja menor que filhos)
    heap.bubbleDown(index)

    return nil

}