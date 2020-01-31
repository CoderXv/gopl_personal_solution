package nodes

type Node struct {
	data int
	next *Node
}

func (node *Node) ChangeData(input int) {
	node.data = input
}

func (node *Node) ToNext(input *Node) {
	node.next = input
}
