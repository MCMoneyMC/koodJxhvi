package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {
	if l1.Tail != nil {
		l1.Tail.Next = l2.Head
	} else {
		l1.Head = l2.Head
	}
}
