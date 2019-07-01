/**
* @Author: YuDC
* @Date: 2019-06-30 23:15
* @Description: time
 */
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// 当前时间
	now := time.Now()
	fmt.Println("now = ", now)

	// 获取年月日 时分秒

	fmt.Printf("年: %d , 月: %d, 日: %d, 时: %d, 分: %d, 秒: %d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 时间格式化

	date := now.Format("2006-01-02")
	fmt.Println(date)

	// 常量

	i := 0
	for {
		i++
		//time.Sleep(time.Second) // 休眠一秒
		time.Sleep(time.Microsecond * 100) // 休眠0.1秒
		fmt.Println(i)
		if i == 1 {
			break
		}
	}

	// 时间戳

	fmt.Println(now.UnixNano()) // 精确到纳秒
	fmt.Println(now.Unix())     // 精确到秒

	var str string
	start := time.Now().Unix()
	for i := 0; i < 100000; i++ {
		str += strconv.Itoa(i)
	}
	end := time.Now().Unix()

	fmt.Printf("执行了%d秒", end-start)

}
