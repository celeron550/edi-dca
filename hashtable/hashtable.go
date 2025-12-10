package main

import ( 
	"errors"
	"fmt"
)

type Map interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Remove(key int)
	Size() int
	LoadFactor() float32
	Init()
}

type Tuple struct {
	key   int
	value string
}

type HashTable struct {
	buckets          [][]Tuple
	elementsInserted int
}

func (table *HashTable) Put(key int, value string) {
	if table.LoadFactor() > 0.75{
		table.increaseBuckets()
	}
	put(table.buckets, key, value)
	table.elementsInserted++
}

func put(buckets [][]Tuple, key int, value string){
	bucket := key % len(buckets) 
	buckets[bucket] = append(buckets[bucket], Tuple{
		key: key, 
		value: value})
	
}

func (table *HashTable) increaseBuckets() { 
	// literalmente a mesma coisa de alocacao dinamica de memoria
	increasedBuckets := make([][]Tuple, 8*len(table.buckets))
	for i:=0; i<len(table.buckets); i++{
		for _, tuple := range table.buckets[i] {
			put(increasedBuckets, tuple.key, tuple.value)
		}
	}
	table.buckets = increasedBuckets 
}

func (table *HashTable) LoadFactor() float32{
	return float32(table.elementsInserted) / float32(len(table.buckets))
}

func (table *HashTable) Get(key int) (string, error){
	bucket := key % len(table.buckets) 
	for _, tuple := range table.buckets[bucket] {
		if tuple.key == key{
			return tuple.value, nil
		}
	}
	return  " ", errors.New("a tabela nao possui essa chave")
}

func (table *HashTable) Remove(key int) error{
	bucket := key % len(table.buckets) 
	index := -1
	for i, tuple := range table.buckets[bucket]{
		if tuple.key == key {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("a tabela nao possui essa chave")
	} else {
		// daria pra usar copy e outras fulerage aq, mas parece q esse eh o melhor jeito
		table.buckets[bucket] = append(table.buckets[bucket][:index], table.buckets[bucket][index+1:]...)
		table.elementsInserted--
		return nil
	}
}

func (table *HashTable) Size() int {
	return table.elementsInserted
}

func (table *HashTable) Init(size int) {
	table.buckets = make([][]Tuple, size)
	table.elementsInserted = 0
}

func main() {

	var table HashTable = HashTable{}
	table.Init(10)

	table.Put(2, "America")
	table.Put(8, "ABC")
	table.Put(6, "Fluminese")

	fmt.Println(table.elementsInserted)
	fmt.Println(table.Get(2))
	fmt.Println(table.Get(8))
	fmt.Println(table.Get(6))
	fmt.Println(table.LoadFactor())
	table.Put(22, "America2")
	table.Put(88, "ABC2")
	table.Put(66, "Fluminese2")
	fmt.Println(table.elementsInserted)
	fmt.Println(table.Get(22))
	fmt.Println(table.Get(88))
	fmt.Println(table.Get(66))
	fmt.Println(table.Get(2))
	fmt.Println(table.Get(8))
	fmt.Println(table.Get(6))
	fmt.Println(table.LoadFactor())
	table.Put(1, "Sao Paulo")
	fmt.Println(table.LoadFactor())
	table.Put(3, "Flamengo")
	fmt.Println(table.LoadFactor())
	table.Put(4, "Vasco")
	fmt.Println(table.LoadFactor())
	fmt.Println(table.Get(22))
	fmt.Println(table.Get(88))
	fmt.Println(table.Get(66))
	fmt.Println(table.Get(2))
	fmt.Println(table.Get(8))
	fmt.Println(table.Get(6))
	fmt.Println(table.Get(1))
	fmt.Println(table.Get(3))
	fmt.Println(table.Get(4))

	table.Remove(4)
	fmt.Println(table.Get(4))
	table.Remove(22)
	fmt.Println(table.Get(22))
	fmt.Println(table.LoadFactor())
	fmt.Println(table.Size())

}

