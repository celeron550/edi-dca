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

	ll.Reverse()
	printLinkedList(ll)

	ll.Set(250, 2)
	printLinkedList(ll)

	ll.Remove(1)
	printLinkedList(ll)

	ll.Pop()
	printLinkedList(ll)
	
	
}