package main

import "fmt"

/**
 为什么用 range ？
 - 美观
 - 同时获得 index 和 value 比 java 和 python 只能获得 value 强大
 */


// 数组是值类型（拷贝）
// [] int 表示切片 [x] int 是数组
// 传递过来的数组是被拷贝过的，因为是值传递，除非传递指针类型 *arr[5] int
// y 【5] int [6] int 是不用类型,会出错，go 是不常用数组的，无法确定类型不方便。常用切片
func printArray(arr [5] int) {
	fmt.Println("----------------")
	for i,v:= range arr {
		fmt.Println(i,v)
	}
	fmt.Println("----------------")
}
func main() {
	//定义5个空数组
	var arr1 [5]int
	//初始化数组
	arr2 := [3]int{0,1,2}
	//定义不初始化长度的数组
	arr3 := [...] int {2,4,6,8,10}
	//定义四行五列的数组
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	// 垃圾遍历
	for i:=0 ; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	// 优雅遍历 go语言提供的 range（范围） 关键字
	for i:= range arr3 {
		fmt.Println(arr3[i])
	}

	// 优雅遍历 i：下标 v：值
	for i,v:= range arr3 {
		fmt.Println(i,v)
	}

	// 优雅遍历  v：值 _代替一个不想用的值，任何地方都可用
	for _,v:= range arr3 {
		fmt.Println(v)
	}

	printArray(arr1)
	//cannot use arr2 (type [3]int) as type [5]int in argument to printArray 【5] int [6] int 是不用类型
	//printArray(arr2)
	printArray(arr3)



}
