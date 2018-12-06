package main

import "fmt"

func printSlice(a []int) {
	fmt.Printf("len = %d , cap = %d \n",len(a),cap(a))
}

func main() {

	/**
	 append
	 */
	var arr = [...] int {0,1,2,3,4,5,6,7,8,9}
	slice := arr [4:9]
	// append 一下 没有超过 cap 的值，还是这个数组的 view
	arr1 := append(slice,10)
	// 又 append 一下，超过了 cap 的值，已经被系统重新分配到其他更大的长度的array了，原数组没人用会被go垃圾回收掉
	arr2 := append(arr1,11)

	fmt.Println(arr) // [0 1 2 3 4 5 6 7 8 10]
	fmt.Println(slice) // [4 5 6 7 8]
	fmt.Println(arr1) // [4 5 6 7 8 10]
	fmt.Println(arr2) // [4 5 6 7 8 10 11]
	fmt.Println(arr) // [0 1 2 3 4 5 6 7 8 10] arr2不是 array 的 view，append 在 array 中无法体现

	/**
	 创建 slice
	 */


	 var s  []int
	 for i := 0; i< 100; i++ {
	 	 // 这里会了解到 切片 cap 的扩充方式 *2 相当于 java 的hashMap
		 printSlice(s)
		 s = append(s, i*2 +1)
	 }

	 s1 := []int{2,4,6,8}
	 //定义一个 10长切片
	 s2 := make([]int,10)
	 //定义一个10长切片 cap 32
	 s3:= make([]int,10,32)
	 printSlice(s1)
	 printSlice(s2)
	 printSlice(s3)

	 /**
	  copy slice
	  */
	  copy(s2,s1)
	  fmt.Println(s2) //[2 4 6 8 0 0 0 0 0 0]

	  /**
	   删除下标4的元素   ...：语法糖展开元素   pop同理 去尾 len(s2) - 1
	   */
	   s2 = append(s2[:3],s2[4:]...)
	   fmt.Println(s2)
}
