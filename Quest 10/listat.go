package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	pos++
	temp := &NodeL{Data: nil, Next: l}
	for temp != nil && pos > 0 {
		temp = temp.Next
		pos--
	}
	return temp
}
