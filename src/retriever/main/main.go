package main

import (
	"fmt"
	"retriever"
	"retriever/httpRetriever"
)

func download(r retriever.Retriever) string {
	return r.Get("http://www.gitrue.com")
}

func main() {
	var r  retriever.Retriever
	r = httpRetriever.HttpRetriever{"name"}

	fmt.Println(download(r))
	// %T ： type
	fmt.Printf("%T , %v\n",r,r)

	/**
	 如何知道接口的类型 type ？
	 */
	//1 - switch type： 已知或者未知的对象数据类型均可，r.(type)必须配合switch来使用，不能单独执行此语句。
	switch v := r.(type) {
	case httpRetriever.HttpRetriever:
		fmt.Println(v.Name)
	}

	// 2 - type assertion (断言）类似 java 中的<泛型> 调用其他实现者的方法会直接报错

	//httpRetriever := r.(httpRetriever.HttpRetriever)
	//fmt.Println(httpRetriever.Name)

	if httpRetriever,ok := r.(httpRetriever.HttpRetriever); ok {
		fmt.Println(httpRetriever.Name)
	} else {
		fmt.Println("not a httpRetriever ")
	}




}
