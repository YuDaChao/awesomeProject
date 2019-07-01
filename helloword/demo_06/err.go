/**
* @Author: YuDC
* @Date: 2019-06-30 14:54
* @Description: 异常捕获
 */
package main

import (
	"errors"
	"fmt"
)

func testErr() {

	// 异常处理

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到错误。。。")
		}
	}()

	n1 := 1
	n2 := 0
	res := n1 / n2
	fmt.Println("res = ", res)
}

// 自定义错误

func testCusError(name string) (err error) {
	if name == "ok" {
		return nil
	}
	return errors.New("发生异常信息")
}

func main() {

	//testErr()

	err := testCusError("error")
	if err != nil {
		// 使用panic终止程序
		panic(err)
	}

	fmt.Println("出现错误之后  ======")
}
