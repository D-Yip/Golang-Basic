package main

import (
	"Golang-Basic/tree"
	"fmt"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Right.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(root)
	fmt.Println(nodes)
	root.Print()
	root.SetValue(5)
	root.Print()

	pRoot := &root
	pRoot.SetValue(1)
	root.Print()
	pRoot.Print()

	pRoot = nil
	pRoot.SetValue(10)

	root.Traverse()
}
