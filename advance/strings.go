/**
* @Author: YuDC
* @Date: 2019-08-01 19:28
* @Description: strings
 */
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// 判断字符串s是否包含字串substr
func testContains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// 判断字符串s是否包含字符串chars中的任意字符
func testContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// 返回字符串s中有几个不重复的sep字串
func testCount(s, sep string) int {
	return strings.Count(s, sep)
}

// 字符串拼接
func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func main() {

	fmt.Println(testContains("hello", "lo")) // true
	fmt.Println(testContains("hello", "el")) // true
	fmt.Println(testContains("hello", ""))   // true
	fmt.Println(testContains("", ""))        // true

	fmt.Println("===============")

	fmt.Println(testContainsAny("hello", "o"))   // true
	fmt.Println(testContainsAny("hello", "x o")) // true
	fmt.Println(testContainsAny("hello", ""))    // false
	fmt.Println(testContainsAny("", ""))         // false

	fmt.Println("===============")

	fmt.Println(testCount("hello", "l")) // 2
	// 当sep为空时，Count的返回值是utf8.RuneCountInString(s) + 1
	fmt.Println(testCount("hello", ""))          // 6
	fmt.Println(utf8.RuneCountInString("hello")) // 5
	fmt.Println(testCount("fivevev", "vev"))     // 1

	// 字符串分割为[]string
	// Fields FieldsFunc  Split SplitAfter  SplitN SplitAfterN

	// Fields 使用一个或多个空格分割字符串
	fmt.Println(strings.Fields("hello golang")) // [hello golang]
	strSlice := strings.FieldsFunc("hello golang", func(r rune) bool {
		return r == 'o'
	})
	fmt.Println(strSlice) // [hell  g lang]

	// 让字符串s按照sep进行分割，不保留sep字串
	fmt.Println(strings.Split("foo,bar,baz", ",")) // [foo bar baz]
	// 让字符串s按照sep进行分割，保留sep字串
	fmt.Println(strings.SplitAfter("foo,bar,baz", ",")) // [foo, bar, baz]
	// 带N的方法
	// N < 0 返回所有字串
	fmt.Println(strings.SplitN("foo,bar,baz", ",", -1)) // [foo bar baz]
	// N == 0 返回[]
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 0)) // nil
	// N > 0 表示返回的slice中最多只有n个元素，其中最后一个元素不会被分割
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 1))  // ["foo,bar,baz"]
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 10)) // ["foo" "bar" "baz"]

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	// 判断字符串是否有某个前缀或后缀
	fmt.Println(strings.HasPrefix("hello", "he"))
	fmt.Println(strings.HasSuffix("hello", "lo"))

	// 字符串或者字串在字符串中出现的位置
	fmt.Println(strings.Index("hello", "e")) // 1
	fmt.Println(strings.Index("hello", "j")) // -1
	// chars中任何一个Unicode在s中首次出现的位置
	fmt.Println(strings.IndexAny("hello", "o l")) // 2
	fmt.Println(strings.IndexRune("hello", 'h'))  // 0

	fmt.Println(Join([]string{"h", "e", "l"}, ","))

	// 字符串重复几次
	fmt.Println(strings.Repeat("ha", 3)) // hahaha

	// 字符串替换
	fmt.Println(strings.Replace("hello", "l", "o", -1)) // heooo

	// 替换多个
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!")) // This is &lt;b&gt;HTML&lt;/b&gt;!

}
