package main

import (
	"fmt"
)

func BubbleSort(v []int) {
	trocou := false
	n := len(v)
	for varredura := 0; varredura < n-1; varredura++ {
		for j := 0; j < (n - varredura - 1); j++ { // Dica 2: sempre vai diminuindo em relação a varredura, para evitar comparação desnecessária
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
				trocou = true
			}
		}
		if !trocou {
			return
		}
	}
}

func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ { //varredura
		menor := i
		for j := i + 1; j < len(v); j++ { //comparação
			if v[j] < v[menor] {
				menor = j
			}
		}
		v[i], v[menor] = v[menor], v[i]
	}
}

// logica:
// vetor ordenado começa pelo primeiro elemento, ai vc vai iterando do final do 
// vetor ate o começo e inserindo na posição correspondente
func InsertionSort(v []int) {
	n := len(v)
	for i := 1; i < n; i++ { // varredura (o elemento 0 está ordenado)
		for j := i; n > 0; j-- { // itera sobre o vetor desordenado (começa pelo final)
            // se tiver desordenado, troca as posi
			if v[j] < v[j-1] { 
				v[j-1], v[j] = v[j], v[j-1]
			} else {
				break
			}
		}
	}
}

func merge(v []int, e []int, d []int) {
	index_e := 0
	index_d := 0
	index_v := 0

    // preenche o vetor V ordenando
	for index_e < len(e) && index_d < len(d) {
		if e[index_e] < d[index_d] {
			v[index_v] = e[index_e]
			index_e++
		} else {
			v[index_v] = d[index_d]
			index_d++
		}
		index_v++
	}
	// agora tem que inserir os elemetos que possam ter sobrado
	for index_e < len(e) {
		v[index_v] = e[index_e]
		index_e++
		index_v++
	}
	for index_d < len(v) {
		v[index_v] = e[index_d]
		index_d++
		index_v++
	}
}

func MergeSort(v []int) {
    if len(v) > 1 {
        // etapa 1: criar os sub-vetores
        mid := len(v)/2
        v_e := make([]int,mid)
        v_d := make([]int,len(v)-mid) // len(v)-mid pra funcionar com vetores impares

        // etapa 2: preencher os subvetores
        index_v := 0
        for ie := 0; ie < len(v_e); ie++{
            v_e[ie] = v[index_v]
            index_v++
        }
        for id := 0; id < len(v_d); id++{
            v_d[id] = v[index_v]
            index_v++
        }

        // etapa 3: chamada recursiva
        MergeSort(v_e)
        MergeSort(v_d)
        Merge(v,v_e,v_d)

    }
}

func QuickSort(v [] int, ini int, fim int){
    if ini < fim {
        index_pivot := Partition(v, ini, fim) // pivotação 
        QuickSort(v,ini,index_pivot-1) // organiza a esqueda
		QuickSort(v,index_pivot+1,fim) // organiza a direita
    }
}

func Partition(v [] int, ini int, fim int) int {
    pivot := v[fim] // define o pivo sendo o ultimo elemento  
    p_index := ini // e ele será o primeiro elemento do vetor
    for i := ini; i < fim; i++ {
        if v[i] <= pivot { //se for maior, troca de posição e incrementa o p_index
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
	fmt.Println("Hello, World!")
	v := []int{2, 8, 6, 10, 4, 5, 3}
	fmt.Println(v)
	// BubbleSort(v)
	// SelectionSort(v)
	InsertionSort(v)
	fmt.Println(v)
}