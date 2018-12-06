package main

import (
	"fmt"
	"tree"
)



func main() {

	var root tree.Node

	root = tree.Node{Value:5}

	root.Left = &tree.Node{}
	root.Right = &tree.Node{5,nil,nil}

	// new 内建函数 返回一个新的地址
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	//fmt.Print(treeNVode{})
	//fmt.Print(root.right.left)
	fmt.Println(root)
	root.Print()

	notes := []tree.Node {
		{Value : 3},
		{},
		{3,nil,&root},
	}

	fmt.Println(notes)

	// 便利树
	root.Traverse()

}