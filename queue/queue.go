package queue

type IQueue interface{
	Push(val int)
	Pop() (int,error)
	Peek() (int,error)
	Size() int
	IsEmpty() bool
}