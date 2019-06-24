package main

import "fmt"

func main() {
	// iota 枚举 只能用于常量
	const (
		a int = iota
		b     = iota
		c     // 相当于 c = iota
		// 遇到const 会变成0
		d int = iota // 0
	)
	fmt.Printf("a = %d, b = %d, c = %d, d = %d \n", a, b, c, d)
	const i, j, k = iota, iota, iota // 0
	fmt.Printf("i = %d, j = %d, k = %d", i, j, k)
}
