package main

import "fmt" // 导入包

func main() {
	// 定义变量, 声明的变量必须要使用
	// 没有赋初值时，系统会根据类型自动赋值
	var a int                     // 默认值为 0
	var b int = 3                 // 赋初始值
	var c = 4                     // 可以省略 类型(int)
	var i, j, k = 1, "jack", true // 可以同时定义多个值
	var (
		x = 1
		y int
		z = 9
	)
	// 定义常量
	const (
		m = iota
		q
		w
	)
	fmt.Printf("a = %d, b = %d, c = %d \n", a, b, c)
	fmt.Printf("i = %d, j = %q, k = %t \n", i, j, k)
	fmt.Printf("x = %d, y = %d, z = %d \n", x, y, z)
	fmt.Println(m, q, w) // 1 1 1

}
