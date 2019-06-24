package main

import "fmt"

func sort(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("第%d次, arr = %v \n", i, arr)
	}
}

func main() {
	// 定义一个数组 长度为10
	// 声明没有赋初值，则根据定义的类型自动赋值
	var arr [10]int
	fmt.Println("arr : ", arr)

	//定义并赋值
	var arr2 = [5]int{5, 1, 2, 4, 3}
	fmt.Println("arr2 : ", arr2)

	// 部分赋值, 没有赋值的默认为初始值
	arr3 := [5]int{1, 2, 3}
	fmt.Println("arr3 : ", arr3)

	// 给指定的位置赋值 下标
	arr4 := [5]int{2: 4, 3: 6}
	fmt.Println("arr4: ", arr4)

	sort(arr2)

	fmt.Println("arr2 ===> ", arr2)

	a := []int{1, 3, 4}
	fmt.Println(a)
}
