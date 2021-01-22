package main

import "fmt"

func main() {
	name := "issone" // 普通变量：存数据值本身
	nameP := &name   // 指针变量：存值的内存地址.  获取name的指针, 取地址.   指针类型只占用 4 个或者 8 个字节的内存大小。
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)

	nameV := *nameP
	fmt.Println("nameP指针指向的值为:", nameV)
	fmt.Printf("nameP 指针类型是：%T\n", nameP)

	*nameP = "issone 123" //修改指针指向的值
	fmt.Println("nameP指针指向的值为:", *nameP)
	fmt.Println("name变量的值为:", name)

	// 创建指针的几种方式
	// 1. 先定义对应的变量，再通过变量取得内存地址，创建指针
	name1 := "string"
	nameP1 := &name1
	fmt.Println("nameP1", nameP1)

	// 2. 先创建指针，分配好内存后，再给指针指向的内存地址写入对应的值。
	// 创建指针
	nameP2 := new(string)
	// 给指针赋值
	*nameP2 = "string2"
	fmt.Println("nameP2", nameP2)

	//

	var nameP3 *string            // 声明一个指针
	fmt.Println("nameP3", nameP3) // 过 var 声明的指针变量是不能直接赋值和取值的，因为这时候它仅仅是个变量，还没有对应的内存地址，它的值是 nil。
	name3 := "string3"
	nameP3 = &name3 // 初始化
	fmt.Println("nameP3", nameP3)
	fmt.Println("*nameP3", *nameP3)

	// 指针使用建议
	// 1. 如果接收者类型是 map、slice、channel 这类引用类型，不使用指针；

	// 2. 如果需要修改方法接收者内部的数据或者状态时，需要使用指针；

	// 3. 如果需要修改参数的值或者内部数据时，也需要使用指针类型的参数；

	// 4. 如果是比较大的结构体，每次参数传递或者调用方法都要内存拷贝，内存占用多，这时候可以考虑使用指针；

	// 5. 像 int、bool 这样的小数据类型没必要使用指针；

	// 6. 如果需要并发安全，则尽可能地不要使用指针，使用指针一定要保证并发安全；

	// 7. 指针最好不要嵌套，也就是不要使用一个指向指针的指针，虽然 Go 语言允许这么做，但是这会使你的代码变得异常复杂。

}
