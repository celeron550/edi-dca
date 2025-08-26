package stack

import "errors"

type Node struct{
	val int
	next *Node
}

type LinkedStack struct{
	topo *Node
	qtd int
}

func (stack *LinkedStack) Size() int {
	return stack.qtd
}

func (stack *LinkedStack) Peek() (int,error){
	if stack.qtd == 0 {
		return -1, errors.New("Pilha vazia")
	}

	return stack.topo.val, nil
}

func (stack *LinkedStack) Push(value int) {
	Novo := &Node{
		val: value,
		next: stack.topo,
	}
	stack.topo = Novo
	stack.qtd++
}

func (stack *LinkedStack) Pop() (int,error){
	if stack.qtd == 0 {
		return -1, errors.New("Pilha vazia")
	}
	aux := stack.topo
	stack.topo = stack.topo.next
	stack.qtd --
	return aux.val, nil
}

func (stack *LinkedStack) IsEmpty() bool {
	if stack.qtd > 0{
		return false
	}
	return true
}
