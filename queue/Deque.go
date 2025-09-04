package queue

type Node2P struct{
	val int
	prev *Node2P
	next *Node2P
}

type Deque struct{
	head     *Node2P
	tail     *Node2P
	inserted int
}

func (d *Deque) Size() int {
	return d.inserted
}

func (d *Deque) Push_Back(e int){

}

func (d *Deque) Push_Front(e int){

}

func (d *Deque) Peek_Back() (int,error){

}

func (d *Deque) Peek_Front() (int,error){

}

func (d *Deque) IsEmpty() bool{

}

