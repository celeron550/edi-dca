package main

import(
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (int,error)
	Add(e int)
	AddOnIndex(e int, index int) error
	Remove(index int) error


}

type ArrayList struct {
	v []int
	inserted int
}

func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}


func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *ArrayList) Get(index int) (int,error) {
	if index >=0 && index < list.inserted {
		return list.v[index], nil
		
	} else {
		return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	}
}

func (list *ArrayList) Add(e int) {
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = e
	list.inserted++
}

func (list *ArrayList) doubleV(){

	newV := make([]int, list.inserted  *2)

	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
		
	}
	list.v = newV
}

func main() {

	l := &ArrayList{}
	l.Init(10)
	for i:=1; i <= 50; i++{
		l.Add(i)
	}
	val, _ := l.Get(0)
	fmt.Println("Valor na posicao 0: ",val)
	
	val, _ = l.Get(49)
	fmt.Println("Valor na posicao 49: ",val)


}
