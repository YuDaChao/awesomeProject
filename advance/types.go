/**
* @Author: YuDC
* @Date: 2019-07-24 20:14
* @Description: 类型
	1: bool
	2: Numeric Types
		1: int8 int16 int32 int64 int
			* int8
				size: 8bits
				range: -128 ~ 127
			* int16
				size: 16bits
				range: -32768 to 32767
			* int32
				size: 32bits
				range: -2147483648 to 2147483647
			* int64
				size: 64bits
				range: -9223372036854775808 to 9223372036854775807
			* int
				size: 32bits 或 64bits
		2: uint8 unit16 unit32 uint64 uint
			* uint8
				size: 8bits
				range: 0 ~ 255
			* uint16
				size: 16bits
				range: 0 to 65535
			* uint32
				size: 32bits
				range: 0 to 4294967295
			* uint64
				size: 64bits
				range: 0 to 18446744073709551615
			* uint
				size: 32bits 或 64bits
		3: float32 float64
		4: complex64 complex128
		5: byte
			byte 其实就是unit8
		6: rune
			rune 其实就是int32
	3: string
*/
package main

import (
	"fmt"
	"unsafe"
)

func testBool() {
	a := true
	b := false
	fmt.Printf("a = %v, b = %v\n", a, b) // a = true, b = false
	c := a && b
	fmt.Println("c = ", c) // c =  false
	d := a || b
	fmt.Println("d = ", d) // d =  true
}

func testInt() {
	var a = 89
	b := 95
	fmt.Println("value of a is ", a, "and b is ", b)                      // value of a is  89 and b is  95
	fmt.Printf("type of a is %T, size of a is %d\n", a, unsafe.Sizeof(a)) // type of a is int, size of a is 8
	fmt.Printf("type of b is %T, size of b is %d\n", b, unsafe.Sizeof(b)) // type of b is int, size of b is 8
}

func testFloat() {
	a, b := 5.67, 8.97                      // 自动推断类型，浮点型默认为float64
	fmt.Printf("type of a %T b %T\n", a, b) // type of a float64 b float64
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff) // sum 14.64 diff -3.3000000000000007

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no2-no1) // sum 145 diff 33
}

func testComplex() {
	// 使用简短方式声明复数
	c := 6 + 7i
	fmt.Println("c", c)

	c1 := complex(5, 7)
	c2 := 8 + 27i
	sum := c1 + c2
	fmt.Println("sum:", sum) // sum: (13+34i)

	mul := c1 * c2
	fmt.Println("mul:", mul) // mul: (-149+191i)
}

func main() {

	testBool()
	testInt()
	testFloat()
	testComplex()
}
