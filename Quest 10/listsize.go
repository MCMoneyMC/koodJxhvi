package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	var count int
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
