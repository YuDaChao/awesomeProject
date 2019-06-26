package main

import "fmt"

func main() {

	// 定义一个map
	//m := map[string]int{}

	// 使用make定义map
	m := make(map[string]int)

	// 赋值
	m = map[string]int{"age": 12, "number": 23}

	fmt.Println("m = ", m)

	// 取值
	val, ok := m["age"]

	fmt.Printf("val: %d, ok: %t", val, ok)

	// 遍历

	for key, val := range m {
		fmt.Printf(" key = %q, val = %d", key, val)
	}

}
