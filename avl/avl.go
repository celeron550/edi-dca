package main

import (
	"fmt"
)

type BstNode struct {
	left   *BstNode
	value  int
	right  *BstNode
	height int
	bf     int
}

type AVL interface {
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
	root.right = newRoot.left
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
	} else if hl > hr {
		root.height = hl + 1
	} else {
		root.height = hr + 1
	}

}

func (root *BstNode) RebalanceLeftLeft() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftNeutral() *BstNode {
	return root.RotRight()
}

func (root *BstNode) RebalanceLeftRight() *BstNode {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *BstNode) RebalanceRightRight() *BstNode {
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightNeutral() *BstNode {
	return root.RotLeft()
}

func (root *BstNode) RebalanceRightLeft() *BstNode {
	root.right = root.right.RotRight()
	return root.RotLeft()

}

func (root *BstNode) Rebalance() *BstNode {
	// checa o fator de balanço e rebalanceia
	if root.bf < -1 { // torta para a esquerda
		if root.left.bf <= 0 { // casos Left-Left e Left-Neutral
			return root.RebalanceLeftLeft()
		} else { // caso Left-Right
			return root.RebalanceLeftRight()
		}
	} else if root.bf > 1 { // torta para a direita
		if root.right.bf >= 0 { // casos Right-Right e Right-Neutral
			return root.RebalanceRightRight()
		} else { // caso Right-Left
			return root.RebalanceRightLeft()
		}
	}
	return root // ja ta balanceada
}

func main() {
	fmt.Println("Hello, World!")
}
