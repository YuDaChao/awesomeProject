/**
* @Author: YuDC
* @Date: 2019-07-14 23:47
* @Description: 运用接口对切片进行排序
 */
package main

import (
	"fmt"
	"sort"
)

// 定义一个结构体

type Hero struct {
	Name string
	Age  int
}

// 定义一个切片来存放Hero

type HeroSlice []Hero

// 让HeroSlice实现排序的接口方法
/**
type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}
*/

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	// 条件为true 就执行 Swap
	return hs[i].Age > hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	s := []int{2, 7, -10, 4, 40}

	// 对切片进行生序
	sort.Ints(s)

	fmt.Println("s = ", s)

	// 初始化几个Hero
	h1 := Hero{"Jack", 23}
	h2 := Hero{"Tom", 21}
	h3 := Hero{"Jy", 10}

	hs := HeroSlice{h1, h2, h3}

	length := hs.Len()
	less := hs.Less(0, 1)
	hs.Swap(0, 1) // 会改变hs (接口是引用传递)
	fmt.Println("len = ", length)
	fmt.Println("less = ", less)

	fmt.Println("排序前: ", hs) // [{Jack 23} {Tom 21} {Jy 10}]

	sort.Sort(hs)

	fmt.Println("排序后: ", hs) // [{Jy 10} {Tom 21} {Jack 23}]
}
