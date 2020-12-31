package main

import (
	"fmt"
	"go_learning/src/oop"
)

// 通过组合的方式去扩充类型
type myTreeNode struct {
	node *oop.TreeNode
}

// 后序遍历TreeNode这棵树
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}
func main() {
	var root oop.TreeNode
	root = oop.TreeNode{Value: 3}
	root.Left = &oop.TreeNode{}
	root.Right = &oop.TreeNode{5, nil, nil}
	root.Right.Left = new(oop.TreeNode)
	root.Left.Right = oop.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	nodes := []oop.TreeNode{
		{
			Value: 3,
		},
		{},
		{
			6, nil, nil,
		},
	}
	fmt.Println(nodes)
	fmt.Println(root)

	// 这边即使这个pRoot是地址也没关系
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	var pRoot2 *oop.TreeNode
	pRoot2.SetValue(300)
	pRoot2 = &root
	pRoot2.SetValue(400)
	pRoot2.Print()

	fmt.Println("------------------------")

	root.Traverse()

	fmt.Println("")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
