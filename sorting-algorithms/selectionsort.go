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

func merge(v []int, e []int, d []int) { //funcao auxiliar para juntar os dois vetores
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
func QuickSort(v [] int, ini int, fim int) {
	if ini < fim {
		index_pivot := Partition(v,ini,fim)
		QuickSort(v,ini,index_pivot-1)
		QuickSort(v,index_pivot+1,fim)
	}
}

func Partition(v [] int, ini int, fim int) int {
    pivot := v[fim]
    p_index := ini
    for i := ini; i < fim; i++ {
        if v[i] <= pivot {
            v[p_index], v[i] = v[i], v[p_index]
            p_index++
        	
        } 
    }
    v[fim],v[p_index] = v[p_index],v[fim]
    return p_index
}

func CountingSort(v []int) []int{
	maior, menor := v[0],v[0]
	// achar o maior e o menor
	for i:=1; i<len(v); i++{
		if v[i] < menor {
			menor = v[i]
		}
		if v[i] > maior {
			maior = v[i]
		}
	}
	// vetor de contagem
	c := make([]int,maior-menor+1) 
	for i:=0; i < len(v); i++ { // contar as ocorrencias de cada valor e atualizar o vetor de contagem
		c[v[i]-menor]++
	}
	
	for i :=0; i<len(c); i++{
		c[i+1] += c[i]
	}

	ord := make([]int, len(v))

	for i:=0; i<len(ord); i++{
		ord[c[v[i]-menor]-1] = v[i]
	}
	
	return ord
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
