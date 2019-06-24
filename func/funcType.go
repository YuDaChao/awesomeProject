package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func calc(a, b int, callback FunType) int {
	return callback(a, b)
}

// 定义一个函数类型

type FunType func(int, int) int

func main() {

	var myFunc FunType

	myFunc = add
	result := myFunc(1, 2)
	fmt.Printf("result = %d\n", result)

	sum := calc(2, 4, add)
	fmt.Printf("sum = %d \n", sum)

	// 匿名函数
	var fun = func() {
		fmt.Printf("sum = %d \n", sum)
	}

	fun()

}
