package main

import "fmt"

func main() {
	//i := 0

	//for { // 死循环
	//	i++
	//	if i == 5 {
	//		fmt.Printf("a = %d", i)
	//		break // 跳出当前循环体
	//	}
	//	fmt.Printf("i = %d \n", i)
	//}

	j := 0
	for {
		j++
		if j == 5 {
			continue // 跳过本次循环
		}

		fmt.Printf("j = %d \n", j)

		if j == 8 {
			break
		}
	}
}
