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
	v [] int
	inserted int
}

func (list* ArrayList) Size() int {
	return list.inserted
}

func (list* ArrayList) Get(index int) (int,error) {
	if index >=0 && index < list.inserted-1 {
		return list.v[index], nil
		
	} return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
}

func (list* ArrayList) Add(e int) {
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = e
	list.inserted++
}

func (list* ArrayList) doubleV(){

	newV := make([]int, list.inserted  *2)

	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
		
	}
}

func main() {

	l = &ArrayList()
	l.Init(50)

}
