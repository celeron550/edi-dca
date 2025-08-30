package list

type IList interface {
	Size() int
	Get(index int) (int,error)
	Set(e int, index int) error // elemento, indice
	Add(e int)
	AddOnIndex(e int, index int) error // elemento, indice
	RemoveOnIndex(index int) error
	Pop() (int,error)

}