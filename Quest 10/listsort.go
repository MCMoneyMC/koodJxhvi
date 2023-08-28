package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	sorted := false
	first := l
	current := l
	for !sorted {
		if current.Next == nil {
			current = first
			sorted = checkSorted(first)
		} else if current.Data > current.Next.Data {
			current.Data, current.Next.Data = current.Next.Data, current.Data
		} else {
			current = current.Next
		}
	}
	return l
}

func checkSorted(l *NodeI) bool {
	current := l
	for current.Next != nil {
		if current.Data > current.Next.Data {
			return false
		}
		current = current.Next
	}
	return true
}
