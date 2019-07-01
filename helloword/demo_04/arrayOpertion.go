/**
* @Author: YuDC
* @Date: 2019-06-28 14:36
* @Description: 数组操作
 */
package main

import (
	"fmt"
)

func main() {

	// 定义并初始化数组
	a1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("arr = %v \n", a1)

	// 数组截取，含头不含尾
	a2 := a1[1:]

	a3 := a1[2:4]

	a4 := []int{7, 8, 9, 10}

	fmt.Printf("a4 = %v \n", a4)

	// 会修a4
	a5 := copy(a4, a1)

	fmt.Printf("a5 = %v \n", a5) // 返回的是赋值成功的个数
	fmt.Printf("a4 = %v \n", a4) // [1 2 3 4]

	fmt.Printf("a2 = %v, s3 = %v \n", a2, a3)

}
