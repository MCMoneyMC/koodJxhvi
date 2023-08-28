package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l.Head != nil {
		NodeLForEachIf(l.Head, f, cond)
	}
}

func NodeLForEachIf(node *NodeL, f func(*NodeL), cond func(*NodeL) bool) {
	if cond(node) {
		f(node)
	}
	if node.Next != nil {
		NodeLForEachIf(node.Next, f, cond)
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
