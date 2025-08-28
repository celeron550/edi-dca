package queue

import(
	"errors"
	"math"
)

type ArrayQueue struct{
	v []int
	front int
	rear int
	// inserted int
}


func (queue *ArrayQueue) Init(size int) {
	queue.v = make([]int, size)
	queue.front = -1
    queue.rear = -1
}

func (queue *ArrayQueue) Size() int{
	return int(math.Abs(float64(queue.rear-queue.front))+1)
}

func (queue *ArrayQueue) Push(val int) {
	if queue.front == -1 && queue.rear == -1 {
		queue.front++
		queue.rear++
	} else {
		queue.rear = (queue.rear+1)%len(queue.v)
	}
	queue.v[queue.rear] = val
}

func (queue *ArrayQueue) Pop() (int,error){
	if queue.front == -1 && queue.rear == -1 {
		return -1, errors.New("Fila vazia")
	} else if queue.front == queue.rear {
		queue.front, queue.rear = -1,-1
		return queue.v[queue.front+1],nil
	}
	
	queue.front++
	return queue.v[queue.front-1], nil
}

func (queue *ArrayQueue) Peek() (int,error) {
	if queue.front == -1 && queue.rear == -1 {
		return -1, errors.New("Fila vazia")
	}
	return queue.v[queue.front], nil
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.front == -1 && queue.rear == -1
}
