package main

import (
	"fmt"
	"github.com/celeron550/edi-dca/list"
)

func main() {

	// ---------------- ArrayList ----------------
	fmt.Println("===== Teste ArrayList =====")
	al := &list.ArrayList{}
	al.Init(2)

	al.Add(10)
	al.Add(20)
	list.PrintArrayList(al)

	al.AddOnIndex(15, 1)
	list.PrintArrayList(al)

	al.Set(25, 2)
	list.PrintArrayList(al)

	al.RemoveOnIndex(1)
	list.PrintArrayList(al)

	al.Pop()
	list.PrintArrayList(al)

	// ---------------- LinkedList ----------------
	fmt.Println("\n===== Teste LinkedList =====")
	ll := &list.LinkedList{}

	ll.Add(100)
	ll.Add(200)
	list.PrintLinkedList(ll)

	ll.AddOnIndex(150, 1)
	list.PrintLinkedList(ll)

	ll.Reversed()
	list.PrintLinkedList(ll)

	ll.Set(250, 2)
	list.PrintLinkedList(ll)

	ll.RemoveOnIndex(1)
	list.PrintLinkedList(ll)

	ll.Pop()
	list.PrintLinkedList(ll)

	// ---------------- DoublyLinkedList ----------------
	fmt.Println("\n===== Teste DoublyLinkedList =====")
	dll := &list.DoublyLinkedList{}
	dll.Add(100)
	dll.Add(200)
	list.PrintDLL(dll)

	dll.AddOnIndex(150, 1)
	list.PrintDLL(dll)

	dll.RemoveOnIndex(1)
	list.PrintDLL(dll)

	dll.Pop()
	list.PrintDLL(dll)

}
