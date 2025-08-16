package main

import(
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (int,error)
	Set(e int, index int)
	Add(e int)
	AddOnIndex(e int, index int) error
	Remove(index int) error
	Pop() error


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
		
	}
	return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	
}

func (list *ArrayList) Add(e int) {
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = e
	list.inserted++
}

func (list *ArrayList) AddOnIndex(e, index int) error {
	if index >=0 && index <= list.inserted {
		if list.inserted == len(list.v) {
			list.doubleV()
		}
		for i := list.inserted; i > index; i-- {
			list.v[i] = list.v[i-1]
		}
		list.v[index] = e
		list.inserted++ 
		return nil
	} 
	if index == 0 {
		return errors.New("Nao eh possivel adicionar no arraylist com indice < 0")
	} 
	return errors.New("Nao eh possivel adicionar no arraylist em uma posição maior > arraylist.size")
		
	
}

func (list *ArrayList) doubleV(){

	newV := make([]int, list.inserted  *2)

	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
		
	}
	list.v = newV
}

func (list *ArrayList) Remove(index int) error {
	if index >=0 && index < list.inserted {
		for i := index; i < list.inserted-1; i++ {
			list.v[i] = list.v[i+1]			
		}
		list.inserted--
		return nil
	} 
	if index < 0 {
		return errors.New("Nao eh possivel remover no arraylist com indice < 0")
	} 
	return errors.New("Nao eh possivel remover no arraylist com indice > arraylist.size")
		
	
}

func (list *ArrayList) Pop() error {
	if list.inserted == 0 {
		return errors.New("ArrayList vazio")
	}
	list.inserted--
	return nil
}

func (list *ArrayList) Set(e int, index int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = e
		return nil
	} 
	if index < 0 {
		return errors.New("Nao eh possivel substituir no arraylist com indice < 0")
	} 
	return errors.New("Nao eh possivel substituir no arraylist com indice > arraylist.size")
		
	
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

	_ = l.AddOnIndex(51,50)
	val, _ = l.Get(50)
	fmt.Println("valor na posicao 50: ",val)

	_ = l.Remove(50)
	val, _ = l.Get(49)
	fmt.Println("valor na posicao 49: ",val)

	_ = l.Pop()
	val, _ = l.Get(48)
	fmt.Println("valor na posicao 48: ",val)

	_ = l.Set(666,48)
	val, _ = l.Get(48)
	fmt.Println("valor na posicao 48 apos a troca: ",val)
	


}
