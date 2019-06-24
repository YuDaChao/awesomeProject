package main

// 导入包后必须要使用
import "fmt"


func variableZeroValue()  {
	var name string
	var count int
	fmt.Printf("%q %d\n", name, count)
	var c int
	c = 10
	fmt.Println(c)
}

func main()  {
	// 调用函数
	variableZeroValue()
}
