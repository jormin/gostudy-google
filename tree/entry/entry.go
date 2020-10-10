package main

import (
	"fmt"
	"go-study/tree"
)

func main() {
	var root tree.Node
	//fmt.Println(root)

	root = tree.Node{
		Value: 3,
		Left:  &tree.Node{},
		Right: &tree.Node{Value: 5},
	}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	//fmt.Println(root)

	//nodes := []TreeNode{
	//	*creareNode(3),
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//root.print()
	//root.setValue(1)
	//root.print()
	//
	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()
	//root.print()

	//var pRoot2 *TreeNode
	//pRoot2.setValue(200)
	//pRoot2 = &root
	//pRoot2.setValue(200)

	fmt.Println("Traverse")
	root.Traverse()
	fmt.Println()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Println(nodeCount)

	myNode := tree.MyNode{
		Node: &root,
	}
	myNode.Traverse()
	fmt.Println()

}
