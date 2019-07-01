/**
* @Author: YuDC
* @Date: 2019-06-27 19:06
* @Description: 练习数组
 */
package main

import "fmt"

func main() {
	// 声明一个数组 var 变量名 [数组长度]T
	// 如果定义的时候没有初始化 那么默认值根据数组类型去判断
	var a1 [5]int
	fmt.Printf("a1 = %v \n", a1) // [0 0 0 0 0]
	fmt.Printf("a1 = %T \n", a1)
	fmt.Printf("a1 len is %d \n", len(a1)) // 5

	// 定义并初始化
	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a2 = %v \n", a2)

	// 初始化指定位置的值(下标从0开始)
	a3 := [5]int{1: 10, 3: 13}
	fmt.Printf("a3 = %v \n", a3) // [0 10 0 13 0]

	// 根据初始化的值来确定数组长度
	a4 := [...]int{1: 2, 8: 13}
	fmt.Printf("a4 = %v \n", a4) // [0 2 0 0 0 0 0 0 13]

	// 数组比较
	fmt.Println(a2 == a3) // false
	//fmt.Println(a2 == a4) // 编译错误 只有数组长度相等的才能比较

	// 数组遍历

	for i := 0; i < len(a4); i++ {
		fmt.Printf("index: %d, value: %d \n", i, a4[i])
	}

	for i, v := range a4 {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}

	// 数组容量 cap()
	fmt.Println(cap(a4))
}
