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
	for varredura := 0; varredura < n-1; varredura++ {
		for j := 0; j < (n - varredura - 1); j++ { // Dica 2: sempre vai diminuindo em relação a varredura, para evitar comparacao desnecessaria
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				trocou = true
			} else {
				trocou = false
			}
		}
		if !trocou {
			return
		}
	}
}

func InsertionSort(v []int) {
	
	n := len(v)
	for varredura := 1; varredura < n; varredura++ {
		for i := varredura; i > 0; i-- {
			if v[i-1] > v[i] {
				v[i], v[i-1] = v[i-1], v[i]
			} else {
				break
			}
		}
	}
}

func merge(v []int, e []int, d []int) {
	iv := 0
	ie := 0
	id := 0
	for ie < len(e) && id < len(d) {
		if e[ie] < d[id] {
			v[iv] = e[ie]
			ie++
		} else {
			v[iv] = d[id]
			id++
		}
		iv++
	}
	for ie < len(e) {
		v[iv] = e[ie]
		ie++
		iv++
	}
	for id < len(d) {
		v[iv] = d[id]
		id++
		iv++	
	}
}

func MergeSort(v []int) {
	if len(v) > 1 {
		mid := len(v)/2
		e := make([]int,mid)
		d := make([]int,len(v)-mid)
		
		iv := 0
		for ie:=0; ie < len(e); ie++{
			e[ie] = v[iv]
			iv++
		}

		for id:=0; id < len(d); id++{
			d[id] = v[iv]
			iv++
		}
		MergeSort(e)
		MergeSort(d)
		merge(v,e,d)
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

	l4 := []int{8, 2, 4, 3, 7, 1, 9, 6, 5}
	InsertionSort(l4)
	fmt.Println(l4)

	v := []int {9,4,3,6,3,2,5,7,1,8}
	// e := []int {3,3,4,6,9}
	// d := []int {1,2,5,7,8}
	//merge(v,e,d)
	fmt.Println(v)
	MergeSort(v)
	fmt.Println(v)
}
