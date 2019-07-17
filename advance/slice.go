/**
* @Author: YuDC
* @Date: 2019-07-17 20:12
* @Description: 深入理解slice
 */
package main

import (
	"fmt"
	"unsafe"
)

// 模拟切片底层数据结构
type slice struct {
	array unsafe.Pointer // 指向所引用的数组指针（unsafe.Pointer可以表示任何可寻址的值的指针）
	len   int            // 长度，当前引用切片的长度
	cap   int            // 容量，当前引用切片的容量（底层数组的元素总数）；在实际开发中，cap一定大于或等于len。否则会导致panic
}

func main() {
	// 定义一个数组
	nums := [3]int{}
	// 给数组的第一个元素赋值
	nums[0] = 1

	n := nums[0]
	n = 2

	fmt.Printf("nums: %v\n", nums) // nums: [1 0 0]
	fmt.Printf("n: %d\n", n)       // n: 2

	arr := [4]int{}

	fmt.Printf("arrs: %v\n", arr)

	// 1: 数组的长度是固定的，他的长度是类型的一部分，
	// 因此[3]int 和 [4]int在类型上是不同的，不能称为"一个东西"
	// 不能相互比较或赋值
	//nums = arr // error

	dnums := nums[:]

	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums)) // dnums: [1 0 0], len: 3, cap: 3

	// len 和 cap不同

	snums := nums[:2]

	// len为所引用元素的个数，cap为所引用的数组的元素总长度
	fmt.Printf("snums: %v, len: %d, cap: %d\n", snums, len(snums), cap(snums)) // snums: [1 0], len: 2, cap: 3

	// 切片的创建
	// 1: var []T 或 []T{}
	// func make([]T, len, cap) []T // 使用make的时候，分配一个数组并返回引用该数组的SLice
	// 根据传入的Slice类型，获取其类型能够申请的最大容量大小
	s := make([]int, 3)

	fmt.Printf("s: %v, len: %d, cap: %d\n", s, len(s), cap(s)) // s: [0 0 0], len: 3, cap: 3

	// 扩容
	// 当使用Slice时，若存储的元素长度不断增长，当条件满足扩容的策略时（切片的len等于cap时），
	// 将会触发自动扩容
	// 计算策略
	// 若新的cap大于2倍老的cap时，则扩容后的容量为2倍老的cap（超出了基准值，就只给你需要的容量大小）
	// 若Slice len小于1024个，在扩容时，增长因子为1（也就是3个变6个）
	// 若Slice len大于1024，早扩容时，增长因子为0.25（原本容量的1/4）
	// 也就是小于 1024 个时，增长 2 倍。大于 1024 个时，增长 1.25 倍

	slice1 := make([]int, 1022, 1022)

	fmt.Printf("len: %d, cap: %d\n", len(slice1), cap(slice1)) // len: 1022, cap: 1022

	slice1 = append(slice1, 1)

	// 此时的len=1023，已经大于了上一次cap1022，即达到扩容条件，由于此时的len小于1024 所以增长 2 倍
	fmt.Printf("len: %d, cap: %d\n", len(slice1), cap(slice1)) // len: 1023, cap: 2048

	slice2 := make([]int, 1025, 1025)

	fmt.Printf("len: %d, cap: %d\n", len(slice2), cap(slice2)) // len: 1025, cap: 1025

	slice2 = append(slice2, 1)

	// 此时的len=1023，已经大于了上一次cap，即达到扩容条件，由于此时的len小于1024 所以增长 2 倍
	fmt.Printf("len: %d, cap: %d\n", len(slice2), cap(slice2)) // len: 1026, cap: 1360

	// 坑

	s1 := [3]int{}

	s1[0] = 1

	fmt.Printf("s1 = %v, len: %d, cap: %d\n", s1, len(s1), cap(s1)) // [1, 0, 0] len: 3 cap: 3

	ss := s1[:2]
	ss[0] = 5

	// 结论
	// 在为扩容前，Slice 指向所引用的 Array， 因此在Slice上的变更，会直接修改原Array
	fmt.Printf("ss: %v, len: %d, cap: %d\n", ss, len(ss), cap(ss))  // [5, 0], len: 2, cap: 3
	fmt.Printf("s1 = %v, len: %d, cap: %d\n", s1, len(s1), cap(s1)) // [5, 0, 0] len: 3 cap: 3

	// 扩容
	// 往 Slice append 元素时，若满足扩容策略，也就是假设插入后，原本数组的容量就超过最大值了
	// 这时候内部就会重新申请一块内存空间，将原本的元素拷贝一份到新的内存空间上。此时其与原本的数组就没有任何关联关系了，再进行修改值也不会变动到原始数组
	ss = append(ss, []int{2, 3}...)
	ss[0] = 100
	fmt.Printf("ss: %v, len: %d, cap: %d\n", ss, len(ss), cap(ss))  // [100, 0, 2, 3], len: 4, cap: 6
	fmt.Printf("s1 = %v, len: %d, cap: %d\n", s1, len(s1), cap(s1)) // [5, 0, 0] len: 3 cap: 3

}
