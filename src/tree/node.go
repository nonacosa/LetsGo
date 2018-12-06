package tree

import "fmt"

type Node struct {
	Value    int
	Left,Right *Node

}

/**
 工厂函数
 */
func CreateNode(value int) *Node {
	// go 可以返回局部变量    值传递一般在栈上，treeNode需要多次使用 go 也可能在堆上再回收掉
	return &Node{Value: value}
}

/**
  [值传递] 为 strut 结构定义方法:语法 func (param Strut) method() {...}
   指针传递方便一些，更加优雅一些
 */

func (node *Node) Print() {
	fmt.Println(node.Value)
}


