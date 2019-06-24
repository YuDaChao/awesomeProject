package main

import "fmt"

func main() {

	// 定义一个数组 并初始化
	arr := [5]int{}
	fmt.Printf("arr = %d, len(arr) = %d, cap(arr) = %d \n", arr, len(arr), cap(arr))

	// 定义一个切片（[]里面为空，或者为...）, 并初始化
	//slice := []int{1,2,3}
	//slice := [...]int{1,2,3}
	// 只定义不初始化
	var slice []int
	// 在切片末尾追加
	slice = append(slice, 11)
	fmt.Printf("slice = %d\n", slice)

	// 自动推导类型
	//s := []int{1, 2, 3, 4}
	// make初始化make(切片, 长度, 容量)
	var s = make([]int, 2, 3)
	//fmt.Println(s[7])
	s = append(s, 1, 2, 3, 4)
	s = s[:3]
	fmt.Printf("s = %d， len: %d, cap = %d \n", s, len(s), cap(s))

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//s1 := array[1:4:6]
	s1 := array[3:]
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println("==============")

	x := []int{1, 2, 3, 4, 5, 6}
	var y = append(x, 0) // 不改变原切片
	fmt.Println(x, y)

	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{0, 0}
	//copy(a, b)
	//fmt.Println("a = ", a) // [0 0 3 4 5 6]
	copy(b, a)
	fmt.Println("b = ", b) // [1 2]

}
