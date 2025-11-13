package main

import (
	"fmt"
	
)


type BstNode struct {
    left *BstNode
    value int
    right *BstNode
    height int
    bf int
}

type AVL interface{
	add(int)
	search(int) bool
	min() (int, error)
	max() (int, error)
	height() int
	preOrder()
	inOrder()
	posOrder()
	levelOrder()
	remove(int)
}

func (root *BstNode) RotRight() *BstNode {
    newRoot := root.left
    root.left = newRoot.right
    newRoot.right = root
	
	root.UpdateProperties()
	newRoot.UpdateProperties()
    
	return newRoot

}

func (root *BstNode) RotLeft() *BstNode {
	newRoot := root.right
	root.right = root.left
	newRoot.left = root	
	// ordem importa
	root.UpdateProperties()
	newRoot.UpdateProperties()
	
	return newRoot
}

func (root *BstNode) UpdateProperties() {
	hl := 0
	hr := 0
	if root.left != nil {
		hl = root.left.height
	}
	if root.right != nil {
		hr = root.right.height
	}
	root.bf = hr - hl
	if root.left == nil && root.right == nil { // noh eh folha 
		root.height = 0
	}else if hl > hr {
        root.height = hl + 1
    } else {
        root.height = hr + 1
    }


}

func (root *BstNode) RebalanceLeftLeft() *BstNode{
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftNeutral() *BstNode{
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftRight() *BstNode{
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *BstNode) RebalanceRightRight() *BstNode{
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightNeutral() *BstNode{
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightLeft() *BstNode{
	root.right = root.right.RotRight()
	return root.RotLeft()

}

func (root *BstNode) Rebalance() *BstNode{
	// checa o fator de balanço e rebalanceia
	if root.bf < 0 { // ela ta torta pra esquerda
		if root.left.bf < 0 { // caso left-left e left-neutral
			return root.RebalanceLeftLeft()
		}else{ // caso left-right
			return root.RebalanceLeftRight()
		}
	}
}

func main() {
	fmt.Println("Hello, World!")
}
