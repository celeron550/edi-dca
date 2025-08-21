package list

import(
	"errors"
	"fmt"
)



type Node struct{
	val int
	next *Node
}

type LinkedList struct {
	head *Node
	inserted int
}

func (list *LinkedList) Size() int {
	return list.inserted
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

func (list *LinkedList) Pop() error {
	if list.inserted == 0 {
		return errors.New("LinkedList vazio")
	}
	return list.Remove(list.inserted-1)
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

func (list *LinkedList) Reverse() {
	// vai do head ate o penultimo e troca os next deles
	aux := list.head
	for i := 0 ; i < list.inserted -1 ; i++ {
		// troca aq
		aux , aux.next = aux.next , aux
	}
	// faz o head da lista apontar pro "ultimo(agr primeiro)" elemento
	list.head = aux

}




