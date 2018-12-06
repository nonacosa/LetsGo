package main

import (
	"fmt"
	"io/ioutil"
)

/**
	每种变量默认不赋值的零值
	bool       -> false
	numbers    -> 0
	string     -> ""
	pointers   -> nil
	slices     -> nil
	maps       -> nil
	channels   -> nil
	functions  -> nil
	interfaces -> nil
 */

func grade(score int) string {
	g := ""
	switch  {
	case score < 0 || score > 100 :
		panic(fmt.Sprintf("Wrong score : %d",score)) // 中断程序的执行
	case score < 60:
		g = "C"
	case score < 80:
		g = "B"
	case score < 90:
		g = "A"
	default:

	}
    return g
}

func main() {
	const filename = "abc.txt"

	// contents, err 生命周期仅在 if block 内
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n",contents)
	}

	fmt.Println(

		grade(50),
		grade(89),
		grade(123),
	)

}
