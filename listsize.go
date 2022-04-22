package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	elt := l.Head
	count := 0
	for elt != nil {
		count++
		elt = elt.Next
	}
	return count
}
