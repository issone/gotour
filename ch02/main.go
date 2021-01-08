package main

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	var i int = 10
	fmt.Println(i)
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ",f64 is", f64)
	var bf bool = false
	var bt bool = true
	fmt.Println("bf is", bf, ",bt is", bt)
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is", s1, ",s2 is", s2)
	fmt.Println("s1+s2=", s1+s2)
	var zi int

	var zf float64

	var zb bool

	var zs string

	fmt.Println(zi, zf, zb, zs)

	pi := &i         // & 可以获取一个变量的地址，也就是指针。
	fmt.Println(*pi) // 要想获得指针 pi 指向的变量值，通过*pi这个表达式即可

	i = 20
	fmt.Println("i的新值是", i)

	const name = "飞雪无情" // 常量

	// const (
	// 	one   = 1
	// 	two   = 2
	// 	three = 3
	// 	four  = 4
	// )

	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)

	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)

	//判断s1的前缀是否是H

	fmt.Println(strings.HasPrefix(s1, "H"))

	//在s1中查找字符串o

	fmt.Println(strings.Index(s1, "o"))

	//把s1全部转为大写

	fmt.Println(strings.ToUpper(s1))

	var a byte = 65
	// 8进制写法: var c byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var c byte = '\x41'    其中 \x 是固定前缀

	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)
	// 或者使用 string 函数
	// fmt.Println("a 的值: ", string(a), " \nb 的值: ", string(b))

	var c byte = 'A'
	var d rune = 'B' // 由于 byte 类型能表示的值是有限，只有 2^8=256 个。所以如果你想表示中文的话，你只能使用 rune 类型
	fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数\n", unsafe.Sizeof(c), unsafe.Sizeof(d))
	var name2 rune = '中'
	fmt.Println("name2 的值: ", string(name2))

	// 不管是 byte 还是 rune ，我都是使用单引号，而没使用双引号,在 Go 中单引号与 双引号并不是等价的,单引号用来表示字符,双引号代表字符串
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s", mystr02) // string 的本质，其实是一个 byte数组
}
