package list

import (
	"errors"
	"fmt"
)

type Node2P struct {
	val      int
	previous *Node2P
	next     *Node2P
}

type DoublyLinkedList struct {
	head     *Node2P
	tail     *Node2P
	inserted int
}

func (list *DoublyLinkedList) Size() int {
	return list.inserted
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		if index <= list.inserted/2 {
			aux := list.head
			for i := 0; i < index; i++ {
				aux = aux.next
			}
			return aux.val, nil
		}
		aux := list.tail
		for i := list.inserted-1; i > index; i-- {
			aux = aux.previous
		}
		return aux.val, nil
	}
	return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
}

func (list *DoublyLinkedList) Add(e int) {

	Novo := &Node2P{
		val:      e,
		next:     nil,
		previous: nil,
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
	Novo.previous = aux
	list.tail = Novo
	list.inserted++
	

}

func (list *DoublyLinkedList) AddOnIndex(e, index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Indice invalido: %d", index))
	}
	
	Novo := &Node2P{
		val: e,
		previous: nil,
		next: nil,
	}

	// inserir no inicio
	if index == 0 {
		Novo.next = list.head
		if list.head != nil {
			list.head.previous = Novo
		}
		list.head = Novo
		if list.tail == nil { // lista tava vazia
			list.tail = Novo
		}
		list.inserted++
		return nil
	}

	// inserir no fim 
	if index == list.inserted-1 {
		Novo.previous = list.tail 
		list.tail.next = Novo
		list.tail = Novo
		list.inserted++
		return nil
	}

	var aux *Node2P 
	// inserir na primeira metade
	if index <= list.inserted/2 {
		aux = list.head
		for i:=0; i<index-1; i++{
			aux = aux.next
		}
		Novo.next = aux
		Novo.previous = aux.previous
		aux.previous.next = Novo
		aux.previous = Novo
		list.inserted++
		return nil
	}
	// inserir na segunda metade
	aux = list.tail
	for i:=list.inserted-1; i>index ; i-- {
		aux = aux.previous
	}
	// rouba o previous do aux e passa pro novo, 
	// next do novo tem q apontar pro aux,
	//  e o previous do auxiliar aponta pro novo
	Novo.previous = aux.previous
	Novo.next = aux
	aux.previous.next = Novo
	aux.previous = Novo
	list.inserted++
	return nil
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
	if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Indice invalido %d",index))
	}

	// remover o primeiro
	if index == 0 {
		if list.inserted == 1 {
			list.head = nil
			list.tail = nil
			list.inserted = 0
			return nil
		}
		list.head = list.head.next
		list.head.previous = nil // ja que é o primeiro elemento ele n aponta pra tras
		list.inserted--
		return nil
	}
	// remover o ultimo
	if index == list.inserted-1 {
		list.tail = list.tail.previous
		list.tail.next = nil // ultimo elemento n aponta pra frente
		list.inserted--
		return nil
	}
	var aux *Node2P
	// remover na primeira metade
	if index <= list.inserted/2 {
		aux = list.head
		for i:=0; i<index ; i++ {
			aux = aux.next
		}
		aux.next = aux.next.next
		aux.next.next.previous = aux.next.previous
		list.inserted--
		return nil
	} 
	// remover na segunda metade
	aux = list.tail
	for i:=list.inserted-1; i>index ; i-- {
			aux = aux.previous
	}
	aux.previous.previous.next = aux.previous.next
	aux.previous = aux.previous.previous
	list.inserted--
	return nil
}
func (list *DoublyLinkedList) Pop() error {
	if list.inserted == 0 {
		return errors.New(fmt.Sprintf("Lista vazia"))
	}
	return list.RemoveOnIndex(list.inserted-1)
}
func PrintDLL(list *DoublyLinkedList) {
    if list.head == nil {
        fmt.Println("Lista vazia")
        return
    }

    aux := list.head
    fmt.Print("Lista: ")
    for aux != nil {
        fmt.Printf("%d", aux.val)
        if aux.next != nil {
            fmt.Print(" <-> ")
        }
        aux = aux.next
    }
    fmt.Println()
}

func (list *DoublyLinkedList) Set(e, index int) error {
if index < 0 || index >= list.inserted {
		return errors.New(fmt.Sprintf("Indice invalido: %d", index))
	}
	return nil
} 