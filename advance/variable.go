/**
* @Author: YuDC
* @Date: 2019-07-24 19:50
* @Description: 变量
 */
package main

import "fmt"

func main() {
	// 声明一个变量如果没有被赋值，那么go将会自动初始化一个值(根据定义的数据类型)
	var age int
	fmt.Println("my age is", age) // my age is  0
	age = 25                      // 赋值
	fmt.Println("my age is", age) // my age is 25

	var name string = "Jack"        // 声明一个变量，并初始化
	fmt.Println("my name is", name) // my name is Jack

	// 如果一个变量有初始值，go将会自动根据变量的值来推断变量的类型
	var gender = "男"                    // 类型推断
	fmt.Println("my gender is", gender) // my gender is 男

	// 同时声明多个变量
	var width, height = 30, 50
	fmt.Printf("width = %d, height = %d\n", width, height) // width = 30, height = 50

	// 同时声明多个不同类型的变量
	var (
		Name    = "Tom"
		Age     = 45
		address string
	)
	fmt.Printf("Name = %s, Age = %d, address = %s\n", Name, Age, address) // Name = Tom, Age = 45, address =

	// 简短声明变量
	// 需要注意的是 简短声明必须初始化所有变量，否侧程序将会抛出一个错误 cannot assign 1 values to 2 variables
	x, y := 1, 2
	//x, y := 1 // error assignment mismatch: 2 variable but 1 values
	fmt.Printf("x = %d, y = %d\n", x, y) // x = 1, y = 2
}
