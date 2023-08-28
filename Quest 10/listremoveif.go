package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	var previous *NodeL = nil
	current := l.Head
	for current != nil {
		if current.Data == data_ref {
			if previous == nil && current.Next != nil {
				l.Head = current.Next
			}
			if current.Next != nil {
				current = replaceLink(previous, current)
			} else {
				if previous == nil {
					l.Head = current.Next
					return
				}
				previous.Next = nil
				current = nil
			}
		} else {
			previous = current
			current = current.Next
		}
	}
}

func replaceLink(previous, current *NodeL) *NodeL {
	if previous != nil {
		previous.Next = current.Next
		return previous
	}
	return current.Next
}
