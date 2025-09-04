package main

func LinearSearch(v []int,e int) int{ //O(n)
	for i:=0; i<len(v); i++{
		if v[i]==e{
			return i
		}
	}
	return -1
}

func BinarySearch(v []int, e int, ini int, fim int) int{
	if ini > fim {return -1}

	mid := (fim+ini)/2
	if v[mid] == e {
		return mid
	} else if v[mid] > e { //ta na primeira metade
		return BinarySearch(v,e,ini,mid-1) // mid-1 pq ja checou o mid

	}else{ // segunda metade
		return BinarySearch(v,e,mid+1,fim)
	}
	
}

func BinarySearch_Iterative(v []int, e int) int{
	ini := 0
	fim := len(v)-1
	for fim > ini{
		mid := (fim+ini)/2
		if v[mid] == e {
			return mid
		} else if v[mid] > e { // ta na primeira metade
			fim = mid-1
		} else { // ta na segunda metade
			ini = mid+1
		}
	}
	return -1
}

func main(){

}