package queue

import "errors"

type Node2P struct {
	val  int
	prev *Node2P
	next *Node2P
}

type Deque struct {
	head     *Node2P
	tail     *Node2P
	inserted int
}

func (d *Deque) Size() int {
	return d.inserted
}

func (d *Deque) Push_Back(e int) {
	Novo := &Node2P{
		val: e,
	}
	if d.inserted == 0 {
		d.head, d.tail = Novo, Novo
	}
	Novo.prev = d.tail
	d.tail = Novo
	d.inserted++
}

func (d *Deque) Push_Front(e int) {
	Novo := &Node2P{
		val: e,
	}
	if d.inserted == 0 {
		d.head, d.tail = Novo, Novo
	}
	Novo.next = d.head
	d.head = Novo
	d.inserted++
}

func (d *Deque) Pop_Front() (int, error) {
	if d.inserted == 0 {
		return -1, errors.New("deque vazio")
	}
	val := d.head.val
	if d.inserted == 1{
		d.head, d.tail = nil,nil
		d.inserted--
		return val, nil
	}
	d.head = d.head.next
	d.head.prev = nil
	d.inserted--
	return val, nil
}

func (d *Deque) Pop_Back() (int, error) {
	if d.inserted == 0 {
		return -1, errors.New("deque vazio")
	}
	val := d.tail.val
	if d.inserted == 1{
		d.head, d.tail = nil,nil
		d.inserted--
		return val, nil
	}
	d.tail = d.tail.prev
	d.tail.next = nil
	d.inserted--
	return val, nil
}

func (d *Deque) Peek_Back() (int, error) {
	if d.inserted == 0 {
		return -1, errors.New("deque vazio")
	}
	return d.tail.val, nil
}

func (d *Deque) Peek_Front() (int, error) {
	if d.inserted == 0 {
		return -1,errors.New("deque vazio")
	}
	return d.head.val, nil
}

func (d *Deque) IsEmpty() bool {
	return d.inserted == 0
}

