package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 != nil {
		NodeIForEach(n2, n1, SortListInsert)
		return n1
	}
	return n2
}

func NodeIForEach(node *NodeI, haystack *NodeI, f func(a *NodeI, b int)) {
	if node == nil {
		return
	}
	f(haystack, node.Data)
	if node.Next != nil {
		NodeIForEach(node.Next, haystack, f)
	}
}

func SortListInsert(l *NodeI, data_ref int) {
	if l == nil {
		return
	}

	var previous *NodeI = nil
	current := l

	for current != nil {
		if data_ref < current.Data {
			node := &NodeI{Data: data_ref, Next: current}
			if previous != nil {
				previous.Next = node
				return
			}
			return
		} else if current.Next == nil {
			node := &NodeI{Data: data_ref}
			current.Next = node
			return
		}
		previous = current
		current = current.Next
	}
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
