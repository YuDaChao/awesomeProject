/**
* @Author: YuDC
* @Date: 2019-08-08 13:36
* @Description: 闭包与普通函数的区别
 */
package main

import "fmt"

func printNum(i int) {
	fmt.Println(i)
}

func testFunc() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v)    // 1 2 3
		defer printNum(v) // 3 2 1
	}
}

func testClosure() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v, &v) // 1 2 3 0xc00009c000
		defer func() {     // 闭包
			fmt.Println(v, &v) // 3 3 3 0xc00009c000
		}()
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return base + i, adder2(base + i)
	}
}

func main() {
	add := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(add(i))
	}
}
