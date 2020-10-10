package tree

type MyNode struct {
	Node *Node
}

func (myNode *MyNode) Traverse() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := MyNode{Node: myNode.Node.Left}
	right := MyNode{Node: myNode.Node.Right}
	left.Traverse()
	right.Traverse()
	myNode.Node.Print()
}
