package tree

import "fmt"

type TreeNode struct {
	Value    int
	Left,Right *TreeNode

}

/**
 工厂函数
 */
func CreateNode(value int) *TreeNode {
	// go 可以返回局部变量    值传递一般在栈上，treeNode需要多次使用 go 也可能在堆上再回收掉
	return &TreeNode{Value: value}
}

/**
  [值传递] 为 strut 结构定义方法:语法 func (param Strut) method() {...}
   指针传递方便一些，更加优雅一些
 */

func (node *TreeNode) Print() {
	fmt.Println(node.Value)
}

/**
 树的遍历 先左再又
 */
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()

}


