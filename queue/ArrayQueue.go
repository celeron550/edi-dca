package queue

import(
	"errors"
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
	if queue.front == -1 && queue.rear == -1 { // fila vazia
		return 0
	}
	if queue.rear >= queue.front { // nao deu a volta
		return queue.rear - queue.front + 1
	}
	return len(queue.v) - queue.front + queue.rear + 1 // deu a volta
}

func (queue *ArrayQueue) doubleV() {
	novoV := make([]int,len(queue.v)*2)
	novoI := 0
	for i := queue.front; novoI <len(queue.v); novoI++  {
		novoV[i] = queue.v[i]
		i = (i+1)%len(queue.v)

	}
	queue.v = novoV
}

func (queue *ArrayQueue) Push(val int) {
	if queue.front == -1 && queue.rear == -1 { // fila vazia
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
	}  
	val := queue.v[queue.front]
	if queue.front == queue.rear { // unico elemento na fila
		queue.front, queue.rear = -1,-1
		return val,nil
	}
	
	// avanço circular
	queue.front = (queue.front + 1) % len(queue.v)
	return val, nil
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
