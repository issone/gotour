package main

import (
	"errors"
	"fmt"
)

/*
函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；
函数名称首字母大写代表公有函数，不同的包也可以调用；
任何一个函数都会从属于某一个包。
*/

// 多个形参的类型是一样的，可以只在最后一个参数后写类型的声明
func sum(a, b int) int {
	return a + b
}

// 多值返回
func sum2(a, b int) (int, error) {

	if a < 0 || b < 0 {

		return 0, errors.New("a或者b不能是负数")

	}

	return a + b, nil

}

// 命名返回参数
func sum3(a, b int) (result int, err error) {

	if a < 0 || b < 0 {

		return 0, errors.New("a或者b不能是负数")

	}

	result = a + b

	err = nil

	return

}

// 多个类型一致的参数,  使用 ...类型，表示一个元素为int类型的切片
func sum4(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func sum5(args ...int) int {
	// 利用 ... 来解序列
	result := sum4(args...)
	return result
}

// 多个类型不一致的参数, 如果你希望传多个参数且这些参数的类型都不一样，可以指定类型为 ...interface{}
func myPrintf(args ...interface{}) {

	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	myPrintf(v1, v2, v3, v4)

	// 匿名函数
	sumx := func(a, b int) int {
		return a + b
	}
	fmt.Println(sumx(1, 2))

	// 闭包
	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	fmt.Println(sum5(1, 2, 3, 4))
}
