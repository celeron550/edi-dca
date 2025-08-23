package list

import(
	"fmt"
	
)

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

	al.RemoveOnIndex(1)
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

	ll.Reversed()
	printLinkedList(ll)

	ll.Set(250, 2)
	printLinkedList(ll)

	ll.RemoveOnIndex(1)
	printLinkedList(ll)

	ll.Pop()
	printLinkedList(ll)
	
		// ---------------- DoublyLinkedList ----------------
	fmt.Println("\n===== Teste DoublyLinkedList =====")
	dll := &DoublyLinkedList{}
	dll.Add(100)
	dll.Add(200)
	PrintDLL(dll)

	dll.AddOnIndex(150, 1)
	PrintDLL(dll)

	dll.RemoveOnIndex(1)
	PrintDLL(dll)

	dll.Pop()
	PrintDLL(dll)

}