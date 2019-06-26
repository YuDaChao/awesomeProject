/**
* @Author: YuDC
* @Date: 2019-06-25 23:58
* @Description: 算术运算符
 */
package main

import "fmt"

func main() {

	// 如果 / 两边都是整数 那么结果将去掉小数部分 保留整数
	fmt.Println(10 / 4) // 2
	var n float32 = 10 / 4
	fmt.Println(n) // 2
	// 如果要保留小数，则需要有浮点数参与运算
	fmt.Println(10 / 4.0) // 2.5

	fmt.Println("=======")

	// 计算公式 a % b = a - a / b * b
	fmt.Println("10 % 3 = ", 10%3)    // 1
	fmt.Println("-10 % 3 = ", -10%3)  // -1
	fmt.Println("10 % -3 = ", 10%-3)  // 1
	fmt.Println("-10 % -3 = ", 10%-3) // 1

	// ++ -- 只能单独使用
	var i = 10
	i++ // ok
	// j := i++ // error
	fmt.Println("i = ", i)

	//if i++ > 0 { // error
	//	fmt.Println("i == ", i)
	//}

	if i > 0 {
		fmt.Println(" i > 0")
	}
}
