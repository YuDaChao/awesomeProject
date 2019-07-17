/**
* @Author: YuDC
* @Date: 2019-07-16 12:54
* @Description: 研究defer和return的关系
 */
package main

import "fmt"

// 技巧
// 使用defer时，可以将做如下改写
// eg: return xxx
// 返回值 = xxx
// 调用defer函数
// 空的return

func f1() (i int) {
	defer func() {
		i++
	}()
	return 0
}

func f2() (i int) {
	n := 5
	defer func() {
		n = n + 5
	}()
	return n
}

func f3() int {
	n := 10
	defer func() {
		n++
	}()
	return 1
}

// 作用域
func f4() {
	for i := 0; i < 5; i++ {
		// 一般不建议这样使用 可能会造成资源泄漏
		defer fmt.Println(i) // 4 3 2 1 0  这也就说明了defer在定义时 参数就已经确定了
	}
}

func f5() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}
	fmt.Println("f5 ends")
	/**
	结果为：
	block ends
	f5 ends
	defer runs
	结论：defer并不是在退出当前代码块的作用域时执行的，defer只会在当前函数和方法返回前被调用
	*/
}

// 传值

type Test struct {
	value int
}

func (t Test) print() {
	fmt.Println("print函数中。。。", t.value) // 0
}

func f6() {
	test := Test{}
	defer test.print() // defer调用时其实会对函数中引用外部参数进行拷贝，所以test.value += 1操作并没有修改被defer捕获的test结构体
	test.value += 1
}

func main() {

	var i = f1()

	fmt.Println("f1函数中的返回值 = ", i) // f1函数中的 i =  1

	var j = f2()

	fmt.Println("f2函数中的返回值 = ", j) // f2函数中的返回值 =  5

	var k = f3()

	fmt.Println("f3函数中的返回值 = ", k) // f3函数中的返回值 =  1

	f4()

	f5()

	f6()

}
