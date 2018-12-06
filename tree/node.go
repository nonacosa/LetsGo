package main

import "fmt"

type treeNode struct {
	value int
	left,right *treeNode

}

/**
 工厂函数
 */
func createNode(value int) *treeNode {
	// go 可以返回局部变量    值传递一般在栈上，treeNode需要多次使用 go 也可能在堆上再回收掉
	return &treeNode{value: value}
}

/**
  [值传递] 为 strut 结构定义方法:语法 func (param Strut) method() {...}
   指针传递方便一些，更加优雅一些
 */

func (node *treeNode) print() {
	fmt.Println(node.value)
}

/**
 树的遍历 先左再又
 */
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}

func main() {

	var root treeNode

	root = treeNode{value:5}

	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}

	// new 内建函数 返回一个新的地址
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//fmt.Print(treeNode{})
	//fmt.Print(root.right.left)
	fmt.Println(root)
	root.print()

	notes := []treeNode {
		{value : 3},
		{},
		{3,nil,&root},
	}

	fmt.Println(notes)

	// 便利树
	root.traverse()

}
