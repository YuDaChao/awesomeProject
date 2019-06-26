package main

import "fmt"

func main() {
	//var a = 10
	a := 9
	switch a {
	case 10:
		fmt.Println("10")
		break // 默认会自动加上break
	case 9:
		fmt.Println("9")
	default:
		fmt.Println("default")
	}

	// 赋初值
	switch i := 10; i {
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("default")
	}

	switch i := 9; i {
	case 9:
		fmt.Println("9")
		//fallthrough // 不跳出 继续往下执行
	case 7, 8:
		fmt.Println("7 , 8")
	default:
		fmt.Println("default")
	}

	j := 90
	switch { // 没有条件
	case j > 90: // 放表达式
		fmt.Println("j > 90")
	case j > 80:
		fmt.Println("j > 85")

	}

}
