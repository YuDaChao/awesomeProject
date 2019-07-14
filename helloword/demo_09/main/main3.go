/**
* @Author: YuDC
* @Date: 2019-07-14 15:57
* @Description:
 */
package main

import "fmt"

type X struct {
	Name string
	Age  int
}

type Y struct {
	Name  string
	Score int
}

type Z struct {
	X
	Y
}

func main() {
	var z Z

	z.Score = 12 // ok

	// 如果嵌套的结构体都用相同的字段名 那么访问时必须明确指定访问的时那个
	// 结构体的变量
	//z.Name = "Jack" // ambiguous selector z.Name

	z.Y.Name = "You"

	fmt.Println(z)
}
