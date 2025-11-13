package main

import (
	"errors"
	"fmt"
)

type BstNode struct {
	val   int
	left  *BstNode
	right *BstNode
}

type bst struct {
	root     *BstNode
	inserted int
}

type BST interface {
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

func createBstNode(val int) *BstNode {
	return &BstNode{
		val: val,
	}
}

func (t *bst) add(val int) {
	if t.root == nil {
		t.root = createBstNode(val)
	} else {
		t.root.AddBstNode(val)
	}

	t.inserted++
}

func (no *BstNode) AddBstNode(val int) {
	if val <= no.val {
		if no.left == nil {
			no.left = createBstNode(val)
		} else {
			no.left.AddBstNode(val)
		}
	} else {
		if no.right == nil {
			no.right = createBstNode(val)
		} else {
			no.right.AddBstNode(val)
		}
	}
}

func (t *bst) search(val int) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.searchBstNode(val)
	}
}

func (no *BstNode) searchBstNode(val int) bool {
	if val == no.val {
		return true
	} else if val < no.val {
		if no.left == nil {
			return false
		} else {
			return no.left.searchBstNode(val)
		}
	} else {
		if no.right == nil {
			return false
		} else {
			return no.right.searchBstNode(val)
		}
	}
}

func (t *bst) min() (int, error) {
	if t.root == nil {
		return -1, errors.New("bst vazia")
	} else {
		return t.root.min(), nil
	}
}

func (no *BstNode) min() int {
	val := no
	for val.left != nil {
		val = val.left
	}
	return val.val
}

func (t *bst) max() (int, error) {
	if t.root == nil {
		return -1, errors.New("bst vazia")
	}
	val := t.root
	for val.right != nil {
		val = val.right
	}
	return val.val, nil

}

func (t *bst) height() int {
	if t.root == nil {
		return 0
	} else {
		return t.root.height()
	}
}

func (no *BstNode) height() int {
	h_BstNodeLeft := 0

	if no.left == nil && no.right == nil {
		return 0
	}

	if no.left != nil {
		h_BstNodeLeft = 1 + no.left.height()
	}

	h_BstNodeRight := 0
	if no.right != nil {
		h_BstNodeRight = 1 + no.right.height()
	}

	if h_BstNodeLeft >= h_BstNodeRight {
		return h_BstNodeLeft
	} else {
		return h_BstNodeRight
	}
}

func (no *BstNode) preOrder() {
	fmt.Println(no.val)
	if no.left != nil {
		no.left.preOrder()
	}
	if no.right != nil {
		no.right.preOrder()
	}
}

func (no *BstNode) inOrder() {
	if no.left != nil {
		no.left.inOrder()
	}
	fmt.Println(no.val)
	if no.right != nil {
		no.right.inOrder()
	}
}

func (no *BstNode) posOrder() {
	if no.left != nil {
		no.left.posOrder()
	}
	if no.right != nil {
		no.right.posOrder()
	}
	fmt.Println(no.val)
}

func (t *bst) remove(val int) error {
	if t.root == nil {
		return errors.New("bst vazia")
	} else {
		t.root.removeBstNode(val)
		t.inserted--
		return nil
	}

}

func (no *BstNode) removeBstNode(val int) *BstNode {
	if val < no.val {
		if no.left != nil {
			no.left = no.left.removeBstNode(val)
		} else {
			return nil

		}
	} else if val > no.val {
		no.right = no.right.removeBstNode(val)
	} else {
		// achemo o noh
		if no.left == nil && no.right == nil {
			return nil
		} else if no.left != nil && no.right == nil {
			// filho a esquerda
			return no.left
		} else if no.left == nil && no.right != nil {
			// filho a direita
			return no.right
		} else {
			// dois filhos, acha o menor da direita e bota la
			min := no.right.min()
			no.val = min
			no.right = no.right.removeBstNode(min)
			return no
		}
	}

}

func main() {
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
	res, err := bst.min()
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
