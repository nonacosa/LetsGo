package main

import "fmt"

/**
 切片是数组的视图  切片有弹性且可以当做array的指针（视图）
 slice 本身没数据是个对array 的view
 */

func upoateSlice(s[] int) {
	s[0] = 100
}
func main() {
	arr := [...] int {0,1,2,3,4,5,6,7,8,9}
	s := arr[2:6]
	fmt.Println(s) //[2 3 4 5] 计算机语言一般半开半闭区间包括开头不包括结尾 参考 substring
	/**
	arr[2:6],arr[2:],arr[:6],arr[:] 都是 slice 不算是值算是array的视图
	 */
	fmt.Println(arr[2:6],arr[2:],arr[:6],arr[:])
	arr1 := arr[2:6]
	upoateSlice(arr1)
	// 不同于数组的拷贝，切片本身被修改了，所在的数组也受到了影响，对多个 slice 修改 源数组会受到多次影响
	fmt.Println(arr1) // [100 3 4 5]
	fmt.Println(arr) // [0 1 100 3 4 5 6 7 8 9]

	/**
	 Reslice : 	slice 进一步 slice 都 view 的同一个 array
	 切片从指定地点向后的length进行视图
	 切片取值可以越切片的界但最终不可越底层array的界，因为有个隐藏的 cap，切片可以向后扩展不能向前扩展
	 */
	 fmt.Printf("Reslice len cap \n")
	 arr1 = arr1[:3]
	 arr1 = arr1[:2]
	 // 当前 arr1 视图长度2 视图到真实arr结尾长度8
	 fmt.Println(arr1,len(arr1),cap(arr1))


}
