package main

import (
	"fmt"
	"math"
)

func SelectionSortOP(v []int) []int {
	ord := make([]int, len(v))
	MAX := math.MaxInt32

	for i := 0; i < len(v); i++ { // varredura
		menor := i
		for j := 0; j < len(v); j++ { // comparar o menor com o resto do vetor
			if v[j] < v[menor] {
				menor = j
			}
		}
		ord[i] = v[menor]
		v[menor] = MAX
	}
	return ord

}

func SelectionSortIP(v []int) {
	for i := 0; i < len(v)-1; i++ { // varredura
		menor := i
		for j := i+1; j < len(v); j++ { // comparar o menor com o resto do vetor
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[i],v[menor] = v[menor],v[i]
		
		
	}
	
}


func main() {
	l := []int{9, 15, 7, 235, 8, 1, 2}
	fmt.Printf("%d", SelectionSortOP(l))
	l2 := []int{9, 15, 7, 235, 8, 1, 2}
	fmt.Println()
	SelectionSortIP(l2)
	fmt.Println(l2)

}