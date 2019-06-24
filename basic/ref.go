package main

import (
	"fmt"
)

func main() {
	var a = 10
	// 打印a的值
	fmt.Printf("a = %d \n", a) // a = 10
	// 打印a的内存地址
	fmt.Printf("&a = %v \n", &a) // &a = 0xc0000a2000

	// 定义指针p
	var p = &a
	fmt.Printf("p = %v \n", p)
	fmt.Printf("&a = %v \n", &a)

	// 给p指向的内存变量赋值
	*p = 123

	fmt.Printf("a = %d, *p = %v \n", a, *p)
}
