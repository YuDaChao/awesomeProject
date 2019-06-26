package main

import "fmt"

func main() {
	// 声明格式 var 变量名 类型， 变量声明必须要使用
	var a int // 默认值为0
	fmt.Println(a)

	// 同时声明多个变量 并赋初始值
	//var b, d int = 1, 2
	b, d := 1, 2
	// 交换两个变量的值
	b, d = d, b
	fmt.Println(b, d)

	// 自动推导类型 必须初始化
	c := 30
	// %T 打印变量类型
	fmt.Printf("c type is %T", c)

	// 声明多个变量
	var (
		i int     = 0
		j float64 = 0.0
	)
	fmt.Printf("i = %d, j = %f ", i, j)

}
