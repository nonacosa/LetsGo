package main

import "fmt"

var (
	aa = 3
	bb = 4
	cc = 5
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a,s)
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue() {
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,s,b)
}

func variablTypeDeduction() {
	var a,b,c,d = 3,4,true,"5"
	var s string = "abc"
	fmt.Println(a,s,b,c,d)
}


func varableShorter() {
	a,b,c,s := 3,4,true,"5"
	b = 4
	fmt.Println(a,b,c,s)
}

func enums() {
	// const 组成枚举
	const (
		cpp = iota //自增
		_ //占位
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota) //iota 参与运算 每次左移10
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variablTypeDeduction()
	varableShorter()
	enums()

}
