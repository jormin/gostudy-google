package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

func (node *Node) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node.Ignored.")
		return
	}
	node.Value = Value
}

func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}
