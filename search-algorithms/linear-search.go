package searchalgorithms

func LinearSearch(v []int,e int) int{ //O(n)
	for i:=0; i<len(v); i++{
		if v[i]==e{
			return i
		}
	}
	return -1
}



