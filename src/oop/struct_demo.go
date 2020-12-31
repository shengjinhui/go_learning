package oop

import (
	"fmt"
)

func (node TreeNode) Print() {
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 工厂函数 => 当做构造器来使用
func CreateNode(value int) *TreeNode {
	// 在go语言里面局部变量地址也可以
	return &TreeNode{Value: value}
}
