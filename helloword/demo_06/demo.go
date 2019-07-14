/**
* @Author: YuDC
* @Date: 2019-07-01 19:30
* @Description: 练习
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printDate() {
	var year int
	var month int
	var day int

	for {
		if year == 0 {
			fmt.Println("请输入年份:")
			_, err := fmt.Scanln(&year)
			if err != nil {
				fmt.Println("年份输入错误", err)
				continue
			}
		} else if month == 0 || month < 1 || month > 12 {
			fmt.Println("请输入月份:")
			_, err := fmt.Scanln(&month)
			if err != nil || month < 1 || month > 12 {
				fmt.Println("月份输入错误", err)
				continue
			}
		} else {
			fmt.Println("请输入日:")
			_, err := fmt.Scanln(&day)
			if err != nil {
				fmt.Println("天输入错误", err)
				continue
			}
			break
		}
	}
	fmt.Printf("您输入的时间为: %d年%d月%d日", year, month, day)
}

func gasNumber() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(100)
	fmt.Println(num)
	for i := 1; i <= 10; i++ {
		fmt.Println("请输入:")
		var n int
		_, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Println("请输入数字")
			continue
		}
		if n == num {
			if i == 1 {
				fmt.Println("你太棒了")
				break
			} else if i <= 3 {
				fmt.Println("你很聪明")
				break
			} else if i <= 9 {
				fmt.Println("一般般")
				break
			} else {
				fmt.Println("你终于猜对了")
				break
			}
		} else if n > num {
			fmt.Println("大了")
		} else {
			fmt.Println("小了")
		}
	}
}

func main() {

	//printDate()
	gasNumber()
}
