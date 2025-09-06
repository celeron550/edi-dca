package deque

import "errors"

type ArrayDeque struct {
	v []int
	front int
	rear int
}

func (deque *ArrayDeque) Init(size int) {
	deque.v = make([]int,size)
	deque.front, deque.rear = -1,-1
}

func (deque *ArrayDeque) Size() int {
	if deque.front == -1 && deque.rear == -1 { // fila vazia
		return 0
	}
	if deque.rear >= deque.front { // nao deu a volta
		return deque.rear - deque.front + 1
	}
	return len(deque.v) - deque.front + deque.rear + 1 // deu a volta
}

func (deque *ArrayDeque) doubleV() {
	n := len(deque.v)
	newV := make([]int, n*2)
	size := deque.Size()
	i := deque.front
	for j := 0; j < size; j++ {
		newV[j] = deque.v[i]
		i = (i + 1) % n
	}
	deque.v = newV
	deque.front = 0
	deque.rear = size - 1
}

func (deque *ArrayDeque) EnqueueFront(value int) {
	if deque.Size() == len(deque.v){
		deque.doubleV()
	}
	if deque.Size() == 0 {
		deque.rear++
		deque.front++
	} else { // ja tem elementos inseridos
		if deque.front - 1 == -1 { // se front esta no 0
        deque.front = len(deque.v)-1 // vai pra ultima pos(circula)
    	} else {
        	deque.front--
    	}

	}

	deque.v[deque.front] = value
}

func (deque *ArrayDeque) EnqueueRear(value int) {
	if deque.Size() == len(deque.v){
		deque.doubleV()
	}
	if deque.Size() == 0 {
		deque.rear++
		deque.front++
	} else {
		deque.rear = (deque.rear+1) % len(deque.v)
	}
	deque.v[deque.rear] = value
}

func (deque *ArrayDeque) DequeueFront() (int, error) {
	if deque.Size() ==0 {
		return -1, errors.New("fila vazia")
	}
	val := deque.v[deque.front]
	if deque.rear == deque.front { // so um elemento na fila
		deque.rear, deque.front = -1, -1
	} else {
		deque.front = (deque.front+1) % len(deque.v)
	}
	return val, nil

}

func (deque *ArrayDeque) DequeueRear() (int, error) {
	if deque.Size() ==0 {
		return -1, errors.New("fila vazia")
	}
	val := deque.v[deque.rear]
	if deque.rear == deque.front { // so um elemento na fila
		deque.rear, deque.front = -1, -1
	} else {
		if deque.rear == 0 { // se estava no começo do array, circula pro final
			deque.rear = len(deque.v) - 1
		} else {
			deque.rear--
		}
	}
	return val, nil
}

func (deque *ArrayDeque) Front() (int, error) {
	if deque.Size() ==0 {
		return -1, errors.New("fila vazia")
	}
	val := deque.v[deque.front]
	return val, nil
}

func (deque *ArrayDeque) Rear() (int, error) {
	if deque.Size() ==0 {
		return -1, errors.New("fila vazia")
	}
	val := deque.v[deque.rear]
	return val, nil
}

func (deque *ArrayDeque) IsEmpty() bool {
	return deque.Size() == 0
}
