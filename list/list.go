package list

type List interface {
	Size() int
	Get(index int) (int,error)
	Set(e int, index int) error // elemento, indice
	Add(e int)
	AddOnIndex(e int, index int) error // elemento, indice
	Remove(index int) error
	Pop() error

}