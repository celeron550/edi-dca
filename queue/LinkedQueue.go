package queue

import "errors"

type Node struct {
	val int
	next *Node
}

type LinkedQueue struct{
	inserted int
	front *Node
	rear *Node
}

func (queue *LinkedQueue) Size() int{
	return queue.inserted
}

func (queue *LinkedQueue) Push(val int) {
	novo := &Node{
		val: val,
		next: nil,
	}
	if queue.inserted == 0 {
		queue.front, queue.rear = novo,novo
		
	}else { 
		queue.rear.next = novo
		queue.rear = novo
	}
	queue.inserted++


}

func (queue *LinkedQueue) Pop() (int,error) {
	if queue.inserted == 0 {
		return -1,errors.New("fila vazia")
	}
	val := queue.front.val
	if queue.inserted == 1 {
		queue.front, queue.rear = nil,nil
	}else {
		queue.front = queue.front.next
	}

	queue.inserted--
	return val,nil
}

func (queue *LinkedQueue) Peek() (int,error) {
 if queue.inserted ==0 {
	return -1, errors.New("fila vazia")
 }
 val := queue.front.val
 return val, nil
 
}

func (queue *LinkedQueue) IsEmpty() bool {
	return queue.inserted == 0
}