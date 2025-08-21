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
		for i := list.inserted; i > index; i-- {
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
		list.inserted ++
		return
	}
	
	aux := list.head

	for aux.next != nil {
		aux = aux.next
	}

	aux.next = Novo
	Novo.previous = aux
	list.inserted++
	

}

func (list *DoublyLinkedList) AddOnIndex(e, index int) error {
	if index >= 0 && index <= list.inserted {
		Novo := &Node2P{
			val: e,
			next: nil,
			previous: nil,
		}
		if index == 0 {
			// novo.previous fica nil, next aponta pro antigo "0" 
			// e o head aponta pro novo Node2P (ou usa o Add())
			Novo.next = list.head
			
			list.head = Novo
			list.inserted++
			return nil
		}
		var aux  *Node2P
		if index <= list.inserted/2 {
		aux = list.head
		for i := 0; i < index-1; i++ {
			aux = aux.next
		}
	} else {
		aux = list.tail
		for i := list.inserted - 1; i > index-1; i-- {
			aux = aux.previous
		}
	}

	}
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
}
