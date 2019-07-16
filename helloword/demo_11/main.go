/**
* @Author: YuDC
* @Date: 2019-07-16 00:20
* @Description: 类型断言
 */
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {

	var a interface{}

	var point = Point{1, 2}

	a = point // ok

	var b Point

	//b = a // error

	// 断言
	// 判断a是否指向Point类型
	b = a.(Point)

	if c, ok := a.(Point); ok {
		fmt.Println("转换成功", c)
	} else {
		fmt.Println("转换失败")
	}

	fmt.Println(a)
	fmt.Println(b)
}
