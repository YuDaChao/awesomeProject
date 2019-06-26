package main

import (
	"fmt"
)

func funText() {
	result := 1 / 1
	fmt.Println("result ", result)
}

func main() {
	// 在main函数执行完之前调用
	// 执行顺序 先进后出
	defer fmt.Println("a")
	defer fmt.Println("c")
	fmt.Println("b")
	//funText() // 出现异常
	defer fmt.Println("x")
}
