package main

import (
	"errors"
	"fmt"
)

type Node struct {
	left   *Node
	value  int
	right  *Node
	height int
	bf     int
}

type AVL struct{
	root *Node
	inserted int
}

type IAVL interface {
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

func (root *Node) RotRight() *Node {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root

	root.UpdateProperties()
	newRoot.UpdateProperties()

	return newRoot

}

func (root *Node) RotLeft() *Node {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root
	// ordem importa
	root.UpdateProperties()
	newRoot.UpdateProperties()

	return newRoot
}

func (root *Node) UpdateProperties() {
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

func (root *Node) RebalanceLeftLeft() *Node {
	return root.RotRight()
}

func (root *Node) RebalanceLeftNeutral() *Node {
	return root.RotRight()
}

func (root *Node) RebalanceLeftRight() *Node {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

func (root *Node) RebalanceRightRight() *Node {
	return root.RotLeft()
}

func (root *Node) RebalanceRightNeutral() *Node {
	return root.RotLeft()
}

func (root *Node) RebalanceRightLeft() *Node {
	root.right = root.right.RotRight()
	return root.RotLeft()

}

func (root *Node) Rebalance() *Node {
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

func createNode(value int) *Node{
	return &Node{
		left: nil,
		right: nil,
		value: value,
		height: 0,
		bf: 0,
	}
}

func (avl *AVL) Add(value int) {
	if avl.root == nil{
		avl.root = createNode(value)
	} else{
		avl.root = avl.root.AddNode(value)
	}

	avl.inserted++
}

func (node *Node) AddNode(value int) *Node{
	if value <= node.value{
		if node.left == nil {
			node.left = createNode(value)
		} else {
			node.left.AddNode(value)
		}
	} else {
		if node.right == nil {
			node.right = createNode(value)
		} else {
			node.right.AddNode(value)
		}
	}

	node.UpdateProperties()
	return node
}


func (avl *AVL) Search(value int) bool{
	if avl.root == nil{
		return false
	} else {
		return avl.root.SearchNode(value)
	}
}

func (node *Node) SearchNode(value int) bool {
	if value == node.value {
		return true
	} else if value < node.value {
		if node.left == nil{
			return false
		} else {
			return node.left.SearchNode(value)
		}
	} else {
		if node.right == nil{
			return false
		} else {
			return node.right.SearchNode(value)
		}
	}
}

func (avl *AVL) Min() (int,error){
	if avl.root == nil{
		return -1, errors.New("AVL vazia")
	} else{
		return avl.root.Min(), nil
	}
}

func (node *Node) Min() int{
	no := node

	for no.left != nil {
		no = no.left
	}

	return no.value

}

func (avl *AVL) Max() (int,error){
	if avl.root == nil{
		return -1, errors.New("AVL vazia")
	} else{
		return avl.root.Max(), nil
	}
}

func (node *Node) Max() int{
	no := node

	for no.right != nil {
		no = no.right
	}

	return no.value

}

func (avl *AVL) Height() int{
	if avl.root == nil {
		return 0
	} else {
		return avl.root.Height()
	}
}

func (no *Node) Height() int {
	hl := 0
	hr := 0
	// arvore vazia
	if no.left == nil && no.right == nil {
		return 0
	}
	
	if no.left != nil{
		hl = 1 + no.left.Height()
	}

	if no.right != nil {
		hr = 1 + no.right.Height()
	}

	if hl <= hr {
		return hl
	}
	return hr

}


func (no *Node) preOrder() {
	fmt.Println(no.value)
	if no.left != nil {
		no.left.preOrder()
	}
	if no.right != nil {
		no.right.preOrder()
	}
}

func (no *Node) inOrder() {
	if no.left != nil {
		no.left.inOrder()
	}
	fmt.Println(no.value)
	if no.right != nil {
		no.right.inOrder()
	}
}

func (no *Node) posOrder() {
	if no.left != nil {
		no.left.posOrder()
	}
	if no.right != nil {
		no.right.posOrder()
	}
	fmt.Println(no.value)
}

func (avl *AVL) PreOrder() {
	if avl.root == nil {
		fmt.Println("Árvore vazia.")
		return
	}
	avl.root.preOrder()
}

func (avl *AVL) InOrder() {
	if avl.root == nil {
		fmt.Println("Árvore vazia.")
		return
	}
	avl.root.inOrder()
}

func (avl *AVL) PosOrder() {
	if avl.root == nil {
		fmt.Println("Árvore vazia.")
		return
	}
	avl.root.posOrder()
}

func (avl *AVL) Remove(value int) error{
	if avl.root == nil {
		return errors.New("AVL vazia")
	} else{
		avl.inserted--
		avl.root = avl.root.RemoveNode(value)
		return nil
	}
}

func (no *Node) RemoveNode(value int) *Node {
	if value < no.value {
		if no.left != nil {
			no.left = no.left.RemoveNode(value)
			return no
		} else {
			return nil
		}
	} else if value > no.value {
		if no.right != nil {
			no.right = no.right.RemoveNode(value)
			return no
		} else {
			return nil
		} 
	} else {
		if no.left == nil && no.right == nil {
			return nil
		} else if no.left != nil && no.right == nil {
			return no.left
		} else if no.left == nil && no.right != nil {
			return no.right
		}else{
			max := no.left.Max()
			no.value = max
			no.left = no.left.RemoveNode(max)
			return no
		}
	}
	no.UpdateProperties()
	return no.Rebalance()
}


func main() {
	fmt.Println("🌳 Testando a implementação da Árvore AVL 🌳")
	fmt.Println("--------------------------------------------")

	// 1. Inicialização e Inserção
	tree := AVL{}
	elements := []int{10, 20, 30, 40, 50, 25}

	fmt.Printf("1. Inserindo elementos: %v\n", elements)
	for _, val := range elements {
		tree.Add(val)
		fmt.Printf("   -> Adicionado %d. Nova raiz: %d (Altura: %d)\n", val, tree.root.value, tree.Height())
	}
	

	fmt.Println("--------------------------------------------")

	// 2. Percursos
	fmt.Println("2. Percursos:")
	fmt.Print("   - Pré-Ordem (Raiz, Esquerda, Direita): ")
	tree.PreOrder()
	fmt.Print("   - Em-Ordem (Esquerda, Raiz, Direita - Deve ser ordenado): ")
	tree.InOrder()
	fmt.Print("   - Pós-Ordem (Esquerda, Direita, Raiz): ")
	tree.PosOrder()

	fmt.Println("--------------------------------------------")

	// 3. Busca
	fmt.Println("3. Busca de elementos:")
	searchValues := []int{40, 15, 20}
	for _, val := range searchValues {
		found := tree.Search(val)
		fmt.Printf("   - O valor %d foi encontrado? %t\n", val, found)
	}

	fmt.Println("--------------------------------------------")

	// 4. Min, Max e Altura
	fmt.Println("4. Propriedades da Árvore:")
	minVal, errMin := tree.Min()
	if errMin == nil {
		fmt.Printf("   - Mínimo: %d\n", minVal)
	} else {
		fmt.Println("   - Erro ao obter mínimo:", errMin)
	}

	maxVal, errMax := tree.Max()
	if errMax == nil {
		fmt.Printf("   - Máximo: %d\n", maxVal)
	} else {
		fmt.Println("   - Erro ao obter máximo:", errMax)
	}

	fmt.Printf("   - Altura da Árvore: %d\n", tree.Height())
	fmt.Printf("   - Total de elementos inseridos: %d\n", tree.inserted)

	fmt.Println("--------------------------------------------")

	// 5. Remoção
	fmt.Println("5. Remoção de elementos e Rebalanceamento:")
	removeVal := 50
	fmt.Printf("   - Removendo o valor %d...\n", removeVal)
	errRemove := tree.Remove(removeVal)
	if errRemove != nil {
		fmt.Println("   - Erro na remoção:", errRemove)
	}

	fmt.Printf("   -> Nova raiz: %d (Altura: %d)\n", tree.root.value, tree.Height())
	fmt.Print("   - Em-Ordem após remoção (deve estar ordenado): ")
	tree.InOrder()
	

	// Testando a remoção de um nó que requer rebalanceamento (depende da estrutura atual)
	removeVal2 := 10
	fmt.Printf("   - Removendo o valor %d...\n", removeVal2)
	errRemove2 := tree.Remove(removeVal2)
	if errRemove2 != nil {
		fmt.Println("   - Erro na remoção:", errRemove2)
	}

	fmt.Printf("   -> Nova raiz: %d (Altura: %d)\n", tree.root.value, tree.Height())
	fmt.Print("   - Em-Ordem após segunda remoção: ")
	tree.InOrder()

	fmt.Println("--------------------------------------------")
	fmt.Println("✅ Testes básicos da AVL concluídos. Verifique as saídas para garantir o balanceamento.")
}