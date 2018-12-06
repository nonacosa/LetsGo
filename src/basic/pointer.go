package main

import "fmt"

/**
 go 的指针不能像C语言一样参与运算，要简单一点
 go 语言参数是值传递还是引用传递呢？java、python 大部分引用传递
    --- go 只有值传递一种方式 传递值或值的指针 显示传递
 */

/**
 交换函数,指针的实践
 */

func swap(a,b *int) {
	*a,*b = *b,*a
}

func main() {
	/**
	 介绍指针
	 */
	var a int = 1
	var pa *int = &a
	//指针只能使用指针方式调用 pa = 3 会抛出异常
	*pa = 3
	fmt.Println(a)

	a,b := 3,4
	swap(&a, &b)
	fmt.Println(a,b)

}
