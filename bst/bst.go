package main

import (
	"fmt"
)

type Node struct{
	val int
	left *Node
	right *Node
}

type bst struct {
	root *Node
	inserted int

}

type BST interface {
	add(int)
    search(int) bool
    min() int
    max() int
    height() int
    preOrder()
    inOrder()
    posOrder()
    levelOrder()
    remove(int)


}

func createNode(val int) *Node{
	return &Node{
		val : val,
	}
}

func (t *bst) add(val int) {
	if t.root == nil {
		t.root = createNode(val)
	} else {
		t.root.AddNode(val)
	}

	t.inserted++
}

func (no *Node) AddNode(val int){
	if val <= no.val {
		if no.left == nil {
			no.left = createNode(val)
		} else {
			no.left.AddNode(val)
		}
	} else {
		if no.right == nil {
			no.right = createNode(val)
		} else {
			no.right.AddNode(val)
		}
	}
}

func (t *bst) search(val int) bool{
	if t.root == nil {
		return false
	} else {
		return t.root.searchNode(val)
	}
}

func (no *Node) searchNode(val int) bool{
	if val == no.val{
		return true
	} else if val < no.val {
		if no.left == nil {
			return false
		} else {
			return no.left.searchNode(val)
		}
	} else {
		if no.right == nil {
			return false
		} else {
			return no.right.searchNode(val)
		}
	}
}


func main(){
	bst := &bst{}

	bst.add(10)
	bst.add(5)
	bst.add(3)
	bst.add(8)

	fmt.Println(bst.search(10))
	fmt.Println(bst.search(5))
	fmt.Println(bst.search(3))
	fmt.Println(bst.search(8))
	fmt.Println(bst.search(20))
}