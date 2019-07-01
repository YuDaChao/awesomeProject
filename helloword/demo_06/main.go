/**
* @Author: YuDC
* @Date: 2019-06-30 12:05
* @Description: 字符串内置函数 strings
 */
package main

import (
	"fmt"
	"strconv"
)

func testLen() {
	// len() 统计字符串的长度 按*字节*统计
	// golang中的编码统一utf-8, (ascii的字符(字母/数字)占一个字节，而汉字占三个字节)

	//str := "hello" // len(str) 5
	str := "hello北" // len(str) 8
	fmt.Println("str len = ", len(str))
}

func testRune() {
	str := "hello 北京"
	strs := []rune(str) // 转换成切片
	for i := 0; i < len(strs); i++ {
		fmt.Printf(" strs[%d] = %c \n", i, strs[i])
	}
}

func testStrconv() {
	// Atoi 数字转字符串
	n, err := strconv.Atoi("123")
	if err == nil {
		fmt.Printf("转化成功, type n = %T \n", n) // int
	} else {
		fmt.Println("转化失败, err = ", err)
	}

	// Itoa 字符串转数字
	str := strconv.Itoa(123)
	fmt.Printf("str = %v, str type = %T \n", str, str)
}

func main() {

	testLen()

	testRune()

	testStrconv()
}
