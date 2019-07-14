/**
* @Author: YuDC
* @Date: 2019-06-27 19:06
* @Description: 练习切片
 */
package main

import "fmt"

func fbSlice(n int) {
	slice := make([]int, n, n)

	slice[0], slice[1] = 1, 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}

	fmt.Println("slice = ", slice)

}

func sliceSort(s []int) []int {

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func sliceFind(s []int, target int) {
	var left, right = 0, len(s) - 1

	for {
		midIndex := (left + right) / 2
		mid := s[midIndex] // 中间数
		if left > right {
			fmt.Printf("抱歉 没有找到")
			break
		}
		if mid == target {
			fmt.Printf("找到了 index = %d", midIndex)
			break
		}

		if target > mid {
			left, right = midIndex+1, len(s)-1
		} else {
			left, right = 0, midIndex-1
		}
	}
}

func main() {

	// 定义切片

	var s1 = []int{1, 2, 3}

	fmt.Println(s1)
	fmt.Printf("s1 type is %T \n", s1) // []int

	// 数组生成切片
	arr := [5]int{11, 22, 33, 44, 55}

	fmt.Printf("arr type is %T \n", arr) // [5]int

	// 切片是对数组一个连续片段的引用
	slice := arr[1:4]

	fmt.Printf("slice = %v \n", slice)
	fmt.Printf("slice type is %T \n", slice) // []int

	fmt.Printf("slice[0]的地址是: %v\n", &slice[0]) // 0xc0000181b8
	fmt.Printf("arr[1]的地址是: %v\n", &arr[1])     // 0xc000098008

	// 修改arr的值
	arr[1] = 100

	// 再次打印slice的值
	fmt.Printf("slice[0] = %d \n", slice[0]) // 100

	// 修改slice的值
	slice[1] = 200

	fmt.Printf("arr[2] = %d \n", arr[2]) // 200

	arr2 := [5]int{1, 2, 3, 4, 5}

	slice2 := arr2[1:]
	fmt.Println("arr2[1] ", &arr2[1])
	fmt.Println("slice2[0] ", &slice2[0])
	fmt.Println("slice2: ", slice2)
	slice3 := append(slice2, 1)

	fmt.Println("slice3: ", slice3)

	fmt.Println("arr2[1] ", &arr2[1])
	fmt.Println("slice2[0] ", &slice2[0])
	fmt.Println("slice3[0] ", &slice3[0])

	slice4 := make([]string, 5)

	fmt.Println("slice4: ", slice4)
	fmt.Println("cap(slice4),", cap(slice4))          // 5
	slice4 = append(slice4, "yu")                     // append 一个元素
	slice4 = append(slice4, "yu", "da")               // append 两个元素
	slice4 = append(slice4, []string{"孙悟空", "唐僧"}...) // append 一个切片
	fmt.Println("cap(slice4),", slice4, cap(slice4))  // 10

	arr4 := [6]int{1, 2, 3, 4, 5, 6}

	slice5 := arr4[1:]

	fmt.Println(slice5)

	str := "hello"

	strSlice := []byte(str) // 字符串转换成byte类型的切片

	strSlice[0] = 'y' // 只支持字符和数字 不支持汉字
	str = string(strSlice)
	fmt.Println(strSlice)
	fmt.Println(str) // yello

	// 支持汉字 将字符串转换成 rune类型的切片

	// rune按字符处理 所以支持汉字
	strSlice1 := []rune(str)

	strSlice1[0] = '北'

	str3 := string(strSlice1)

	fmt.Println(str3) // 北ello

	fbSlice(9)

	ss := []int{2, 5, 1, 7, 21, 10}

	res := sliceSort(ss)
	fmt.Println(res)

	ss2 := []int{1, 2, 3, 4, 5, 6}

	sliceFind(ss2, 11)

}
