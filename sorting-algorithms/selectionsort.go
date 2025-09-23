package main

import (
	"fmt"
	"math"
)

func SelectionSortOP(v []int) []int {
	ord := make([]int, len(v))
	MAX := math.MaxInt32

	for varredura := 0; varredura < len(v); varredura++ { // varredura
		menor := varredura
		for j := 0; j < len(v); j++ { // comparar o menor com o resto do vetor
			if v[j] < v[menor] {
				menor = j
			}
		}
		ord[varredura] = v[menor]
		v[menor] = MAX
	}
	return ord

}

func SelectionSortIP(v []int) {
	for varredura := 0; varredura < len(v)-1; varredura++ { // varredura
		menor := varredura
		for j := varredura + 1; j < len(v); j++ { // comparar o menor com o resto do vetor
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[varredura], v[menor] = v[menor], v[varredura]

	}

}

func BubbleSort(v []int) {
	trocou := false
	n := len(v)
	for varredura:=0; varredura<n-1; varredura++ {
		for j:=0; j<(n-varredura-1); j++{ // sempre vai diminuindo em relação a varredura, para evitar comparacao desnecessaria
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1],v[j]
				trocou = true
			}else{
				trocou = false
			}
		}
		if !trocou{break}
	}
}

func main() {
	l := []int{9, 15, 7, 235, 8, 1, 2}
	fmt.Printf("%d", SelectionSortOP(l))
	l2 := []int{9, 15, 7, 235, 8, 1, 2}
	fmt.Println()
	SelectionSortIP(l2)
	fmt.Println(l2)

	l3 := []int{2, 8, 6, 10, 4, 5, 3}
	BubbleSort(l3)
	fmt.Println(l3)

}
