/**
* @Author: YuDC
* @Date: 2019-07-06 16:46
* @Description: map
 */
package main

import "fmt"

func main() {

	// 声明一个map

	var m map[string]string // 此时的m并没有给它分配空间

	//m["name"] = "猪八戒" // 会报错 panic: assignment to entry in nil map

	fmt.Println("m = ", m)

	// 声明并初始化
	//m1 := map[string]string{}
	m1 := map[string]string{"addr": "北京"}

	m1["name"] = "猪八戒"

	fmt.Println("m1 = ", m1)

	// 使用make初始化
	m2 := make(map[string]string)

	m2["name"] = "唐僧"
	m2["name1"] = "唐僧"
	m2["name2"] = "唐僧"

	fmt.Printf("m2 = %v, len(m2) = %d\n", m2, len(m2)) // 3

	// 删除
	delete(m2, "name")

	fmt.Printf("m2 = %v, len(m2) = %d\n", m2, len(m2)) // 3

	mm := make(map[int]int)

	mm[1] = 1
	mm[3] = 3
	mm[4] = 4
	mm[2] = 2

	fmt.Println(mm) // [1:1 2:2 3:3 4:4] 有序的

	mm2 := make(map[string]string)

	mm2["age"] = "23"
	mm2["name"] = "猪八戒"
	mm2["addr"] = "北京"

	fmt.Println(mm2) // map[addr:北京 age:23 name:猪八戒] 也是有序的

}
