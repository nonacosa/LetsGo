package main

import "fmt"

func main() {
	/**
	 创建map 语法 map[key]value {...} 支持复合map
	 */
	m := map[string]string {
		"name":"zwd",
		"tel":"110",
	}
	// 定义 空 map 的区别
	m2 := make(map[string]string) // empty map
	var m3  map[string]string // nil
	fmt.Println(m,m2,m3)

	// map 是 hashMap 无序的
	for k,v := range m {
		fmt.Println(k,v)
	}

	tel := m["tel"]
	telErr,ok := m["tel "]
	fmt.Printf("tel : %s \n",tel)
	// 拼错了 key 不会异常会打印一个空的,多加一个接受参数进行判断
	fmt.Printf("telErr : %s , contanis ? %t \n",telErr,ok)
	// tel,ok tel与上面 tel 不冲突 因为两个参数上面一个参数
	if tel,ok := m["tel"]; ok {
		fmt.Printf("tel : %s \n",tel)
	}

	delete(m,"tel")

}
