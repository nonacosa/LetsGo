package main

import (
	"fmt"
	"tree"
)

/**
 如何扩展已有类型
 */
type MyNode struct {
	node *tree.Node
}

func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	//cannot call pointer method on MyNode literal 不能这样直接用包装过得调用方法，要用变量
	//MyNode{myNode.node.Left}.postOrder()
	left := MyNode{myNode.node.Left}
	right := MyNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}


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

	// cannot call pointer method on MyNode literal
	//MyNode{&root}.postOrder()
	myNote := MyNode{&root}
	myNote.postOrder()


}