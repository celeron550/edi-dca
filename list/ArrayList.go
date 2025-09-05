package list

import(
	"errors"
	"fmt"
)



type ArrayList struct {
	v []int
	inserted int
}





func (l *ArrayList) Init(size int) {
	l.v = make([]int, size)
}


func (list *ArrayList) Size() int {
	return list.inserted
}





func (list *ArrayList) Get(index int) (int,error) {
	if index >=0 && index < list.inserted {
		return list.v[index], nil
		
	}
	return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	
}





func (list *ArrayList) Add(e int) {
	if list.inserted == len(list.v) {
		list.doubleV()
	}
	list.v[list.inserted] = e
	list.inserted++
}



func (list *ArrayList) AddOnIndex(e, index int) error {
	if index >=0 && index <= list.inserted {
		if list.inserted == len(list.v) {
			list.doubleV()
		}
		for i := list.inserted; i > index; i-- {
			list.v[i] = list.v[i-1]
		}
		list.v[index] = e
		list.inserted++ 
		return nil
	} 
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}


func (list *ArrayList) doubleV(){
	newSize := len(list.v) * 2
	if newSize == 0 {
		newSize = 1
	}
	newV := make([]int, newSize)
	for i := 0; i < len(list.v); i++ {
		newV[i] = list.v[i]
		
	}
	list.v = newV
}

func (list *ArrayList) RemoveOnIndex(index int) error {
	if index >=0 && index < list.inserted {
		for i := index; i < list.inserted-1; i++ {
			list.v[i] = list.v[i+1]			
		}
		list.inserted--
		return nil
	} 
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}



func (list *ArrayList) Pop() (int,error) {
	if list.inserted == 0 {
		return -1,errors.New("ArrayList vazio")
	}
	val := list.v[list.inserted-1]
	list.inserted--
	return val,nil
}



func (list *ArrayList) Set(e int, index int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = e
		return nil
	} 
	return errors.New(fmt.Sprintf("Indice invalido: %d", index))
		
	
}

func (list *ArrayList) Reverse() {
	if list.inserted == 0 {
		fmt.Println("Lista vazia")
		return
	}
	for i:=0; i < list.inserted/2; i++ {
		list.v[i],list.v[list.inserted-i-1] = list.v[list.inserted-i-1], list.v[i]
	}
}

func PrintArrayList(list *ArrayList) {
	fmt.Print("ArrayList: [")
	for i := 0; i < list.Size(); i++ {
		val, _ := list.Get(i)
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}






