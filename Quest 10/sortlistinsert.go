package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return &NodeI{Data: data_ref}
	}

	var previous *NodeI = nil
	current := l

	for current != nil {
		if data_ref < current.Data {
			node := &NodeI{Data: data_ref, Next: current}
			if previous != nil {
				previous.Next = node
				return l
			}
			return node
		} else if current.Next == nil {
			node := &NodeI{Data: data_ref}
			current.Next = node
			return l
		}
		previous = current
		current = current.Next
	}
	return l
}
