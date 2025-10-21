package main

import (
	"fmt"
	"errors"
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
    min() (int,error)
    max() (int,error)
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


func (t *bst) min() (int,error) {
	if t.root == nil {
		return -1,errors.New("bst vazia")
	}else {
		val := t.root
		for val.left != nil {
			val = val.left
		}
		return val.val,nil
	}
}

func (t *bst) max() (int,error) {
	if t.root == nil {
		return -1,errors.New("bst vazia")
	}
	val := t.root
	for val.right != nil {
		val = val.right
	}
	return val.val,nil
	
}

func (t *bst) height() int {
	if t.root == nil {
		return 0
	}else {
		return t.root.height()
	}
}

func (no *Node) height() int {
	h_NodeLeft :=0 
	
	if no.left == nil && no.right==nil{
		return 0
	}
	
	if no.left != nil {
		h_NodeLeft = 1+no.left.height()
	}
	
	h_NodeRight := 0
	if no.right != nil {
		h_NodeRight = 1+no.right.height()
	}

	if h_NodeLeft >= h_NodeRight {
		return h_NodeLeft
	} else {
		return h_NodeRight
	}
}

func (no *Node) preOrder(){
	fmt.Println(no.val)
	if no.left != nil {
		no.left.preOrder()
	}
	if no.right != nil {
		no.right.preOrder()
	}
}

func (no *Node) inOrder(){
	if no.left != nil {
		no.left.inOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.inOrder()
	}
}


func (no *Node) posOrder(){
	if no.left != nil {
		no.left.posOrder()
	}
	if no.right != nil {
		no.right.posOrder()
	}
	fmt.Println(no.val)
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
	res , err := bst.min()
	if err != nil {
		return 
	}
	fmt.Println(res)
	max_res, err := bst.max()
	if err != nil {
		return 
	}
	fmt.Println(max_res)
	fmt.Println(bst.height())

	
}
