package main

import(
	"errors"
	"fmt"
)

type List interface {
	Size() int
	Get(index int) (int,error)
	Set(e int, index int) error // elemento, indice
	Add(e int)
	AddOnIndex(e int, index int) error // elemento, indice
	Remove(index int) error
	Pop() error


}

type ArrayList struct {
	v []int
	inserted int
}

type Node struct{
	val int
	next *Node
}

type LinkedList struct {
	head *Node
	inserted int
}


func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}


func (list *ArrayList) Size() int {
	return list.inserted
}

func (list *LinkedList) Size() int {
	return list.inserted
}

func (list *ArrayList) Get(index int) (int,error) {
	if index >=0 && index < list.inserted {
		return list.v[index], nil
		
	}
	return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	
}

func (list *LinkedList) Get(index int) (int, error) {
	if index >=0 && index < list.inserted {
		aux := list.head
		for i := 0; i<index; i++ {
			aux = aux.next
		}
		return aux.val, nil
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

func (list *LinkedList) Add(e int) {

	// cria novo node, fazendo seu next apontar pra nada, depois percorre a lista ate 
	// chegar no ultimo elemento e faz seu next apontar pro novo node

	Novo := &Node{
		val : e,
		next : nil,
	}

	if list.head == nil {
		list.head = Novo
		list.inserted++
		return
	}
	aux := list.head
	
	for aux.next != nil { 
		aux = aux.next
	}
	aux.next = Novo
	list.inserted ++
	return
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
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}

func (list *LinkedList) AddOnIndex(e, index int) error {
	if index >=0 && index <= list.inserted {
		Novo :=&Node{
			val : e,
			next : nil,
		}

		if index == 0 {
			// inserção no começo
			Novo.next = list.head
			list.head = Novo
			list.inserted++
			return nil
		}
		// inserção no meio
		aux := list.head
		for i := 0; i<index-1 ; i++{
			aux = aux.next
		}
		Novo.next = aux.next
		aux.next = Novo

		list.inserted++ 
		return nil
	} 
	
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))

}

func (list *ArrayList) doubleV(){
	newSize := len(list.v) * 2
	if newSize == 0 {
		newSize = 1
	}
	newV := make([]int, newSize)
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
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}

func (list *LinkedList) Remove(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Indice invalido: %d", index))
	}

	if index == 0 {
		// remover o primeiro nó
		list.head = list.head.next
		list.inserted--
		return nil
	}

	// encontrar nó anterior
	anterior := list.head
	for i := 0; i < index-1; i++ {
		anterior = anterior.next
	}

	// pular o nó que será removido
	anterior.next = anterior.next.next

	list.inserted--
	return nil
}

func (list *ArrayList) Pop() error {
	if list.inserted == 0 {
		return errors.New("ArrayList vazio")
	}
	list.inserted--
	return nil
}

func (list *LinkedList) Pop() error {
	if list.inserted == 0 {
		return errors.New("LinkedList vazio")
	}
	return list.Remove(list.inserted-1)
}

func (list *ArrayList) Set(e int, index int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = e
		return nil
	} 
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}

func (list *LinkedList) Set(e, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Indice invalido: %d", index))
	}

	aux := list.head
	for i := 0; i < index; i++ {
		aux = aux.next
	}
	aux.val = e
	return nil
}

func printArrayList(list *ArrayList) {
	fmt.Print("ArrayList: [")
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func printLinkedList(list *LinkedList) {
	fmt.Print("LinkedList: [")
	aux := list.head
	for aux != nil {
		fmt.Print(aux.val)
		if aux.next != nil {
			fmt.Print(", ")
		}
		aux = aux.next
	}
	fmt.Println("]")
}

func main() {

	// ---------------- ArrayList ----------------
	fmt.Println("===== Teste ArrayList =====")
	al := &ArrayList{}
	al.Init(2)

	al.Add(10)
	al.Add(20)
	printArrayList(al)

	al.AddOnIndex(15, 1)
	printArrayList(al)

	al.Set(25, 2)
	printArrayList(al)

	al.Remove(1)
	printArrayList(al)

	al.Pop()
	printArrayList(al)

	// ---------------- LinkedList ----------------
	fmt.Println("\n===== Teste LinkedList =====")
	ll := &LinkedList{}

	ll.Add(100)
	ll.Add(200)
	printLinkedList(ll)

	ll.AddOnIndex(150, 1)
	printLinkedList(ll)

	ll.Set(250, 2)
	printLinkedList(ll)

	ll.Remove(1)
	printLinkedList(ll)

	ll.Pop()
	printLinkedList(ll)
	
}
