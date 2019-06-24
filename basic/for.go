package main

import "fmt"

func main() {
	var str = "hello golang"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%q, \n", str[i])
	}

	for i, data := range str {
		fmt.Printf("i = %d, data = %q \n", i, data)
	}

	for i, _ := range str { // 丢弃第二个返回值
		fmt.Printf("i = %d, data = %q \n", i, str[i])
	}
}
