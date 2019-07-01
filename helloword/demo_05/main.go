/**
* @Author: YuDC
* @Date: 2019-06-29 10:39
* @Description: 函数
 */
package main

import (
	"awesomeProject/helloword/demo_05/util"
	"fmt"
)

// 在main函数执行之前被调用
func init() {
	fmt.Println("init...")
}

func cal(n1 float32, n2 float32, operator byte) float32 {
	var result float32

	switch operator {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		result = n1 / n2
	default:
		result = 0.0
	}
	return result
}

// 闭包
func addUpper() func(int) int {
	var base = 10
	return func(x int) int {
		base++
		return base + x
	}
}

// defer

func testDefer(n1 int, n2 int) int {
	// defer 入栈 后进先出
	defer fmt.Println("defer -----> n1 ", n1) // 1
	defer fmt.Println("defer -----> n2", n2)  // 2
	n1++                                      // 2
	n2++                                      // 3
	result := n1 + n2
	fmt.Println("-----> n1 + n3 = ", result) // 5
	return result
}

func printSj(num int) {
	for i := 1; i <= num; i++ {
		for k := num; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func printCf(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d   ", i, j, i*j)
		}
		fmt.Println("")
	}
}

func main() {

	var n1 float32 = 1.3
	var n2 float32 = 1.7
	var operator byte = '*'
	result := cal(n1, n2, operator)
	fmt.Printf("result = %.2f \n", result)

	// 调用 包名.xxx
	sum := util.CalAdd(n1, n2)

	util.Test()

	fmt.Printf("sum = %.1f \n", sum)

	// 匿名函数
	res := func() int {
		return 10
	}()

	fmt.Println("res = ", res)

	// 闭包
	add := addUpper()
	fmt.Println("add = ", add(1)) // 12
	fmt.Println("add = ", add(2)) // 14
	fmt.Println("add = ", add(3)) // 16

	num := testDefer(1, 2)
	fmt.Println("num = ", num)

	printSj(6)
	printCf(9)
}
