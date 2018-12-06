package main

import "fmt"
import "tree"



func main() {

	var root tree.TreeNode

	root = tree.TreeNode{Value:5}

	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5,nil,nil}

	// new 内建函数 返回一个新的地址
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	//fmt.Print(treeNVode{})
	//fmt.Print(root.right.left)
	fmt.Println(root)
	root.Print()

	notes := []tree.TreeNode {
		{Value : 3},
		{},
		{3,nil,&root},
	}

	fmt.Println(notes)

	// 便利树
	root.Traverse()

}