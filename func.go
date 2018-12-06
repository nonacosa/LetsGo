package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
 跟js差不多：匿名函数 函数式编程 局部函数都可以
 */

func apply(op func(int,int) int ,a,b int) int {
	// omitted 可以节省 := 这里为了熟悉语法先不简写
	var pointer uintptr = reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args %d %d \n",opName,a,b)
	return op(a,b)
}

//重写dim函数
func dim(a,b int) int {
	return int( math.Dim(float64(a),float64(b)))
}


func main() {


	//自定义匿名函数
	pow  := apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i),float64(i2)))
	},2,10)

	//调用已有函数
	dim := apply(dim,10,2)
	fmt.Println(pow)
	fmt.Println(dim)
}
