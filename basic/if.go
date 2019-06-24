package main

import "fmt"

func main() {
	var a = 4
	if a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a < 5")
	}

	if a = 4; a > 5 {
		fmt.Println("a > 5")
	} else {
		fmt.Println("a < 5")
	}
}
