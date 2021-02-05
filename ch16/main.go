package main

import (
	"fmt"
	"unsafe"
)

// unsafe 包里的功能虽然不安全，但的确很香，比如指针运算、类型转换等，都可以帮助我们提高性能

// unsafe 包里最常用的就是 Pointer 指针，通过它可以让你在 *T、uintptr 及 Pointer 三者间转换，从而实现自己的需求，比如零内存拷贝或通过 uintptr 进行指针运算，这些都可以提高程序效率,
// 比如 []byte 转 string，就可以通过 unsafe.Pointer 实现零内存拷贝

// unsafe.Pointer 主要用于指针类型的转换，而且是各个指针类型转换的桥梁。uintptr 主要用于指针运算，尤其是通过偏移量定位不同的内存

// 任何类型的 *T 都可以转换为 unsafe.Pointer；

// unsafe.Pointer 也可以转换为任何类型的 *T；

// unsafe.Pointer 可以转换为 uintptr；

// uintptr 也可以转换为 unsafe.Pointer。

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	// unsafe.Pointer 是一种特殊意义的指针，可以表示任意类型的地址，类似 C 语言里的 void* 指针，是全能型的
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)
	// fmt.Println(*fp)	//  输出打印的时候，不要使用*fp取值，因为类型发生了变化，不知道会是什么，unsafe.Pointer的作用在于指针类型互转的桥梁。
}

func test2() {
	// 声明一个 *person 类型的指针变量 p
	p := new(person)

	// 把 *person 类型的指针变量 p 通过 unsafe.Pointer，转换为 *string 类型的指针变量 pName,
	pName := (*string)(unsafe.Pointer(p))

	// 因为 person 这个结构体的第一个字段就是 string 类型的 Name，所以 pName 这个指针就指向 Name 字段（偏移为 0），对 pName 进行修改其实就是修改字段 Name 的值
	*pName = "isson"

	// 因为 Age 字段不是 person 的第一个字段，要修改它必须要进行指针偏移运算。所以需要先把指针变量 p 通过 unsafe.Pointer 转换为 uintptr，这样才能进行地址运算
	// 偏移量可以通过函数 unsafe.Offsetof 计算出来，该函数返回的是一个 uintptr 类型的偏移量，有了这个偏移量就可以通过 + 号运算符获得正确的 Age 字段的内存地址了，也就是通过 unsafe.Pointer 转换后的 *int 类型的指针变量 pAge
	// 如果要进行指针运算，要先通过 unsafe.Pointer 转换为 uintptr 类型的指针。指针运算完毕后，还要通过 unsafe.Pointer 转换为真实的指针类型（比如示例中的 *int 类型），这样可以对这块内存进行赋值或取值操作
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
}

func test3() {
	// Sizeof 函数可以返回一个类型所占用的内存大小，这个大小只与类型有关，和类型对应的变量存储的内容大小无关，比如 bool 型占用一个字节、int8 也占用一个字节
	fmt.Println(unsafe.Sizeof(true))

	fmt.Println(unsafe.Sizeof(int8(0)))

	fmt.Println(unsafe.Sizeof(int16(10)))

	fmt.Println(unsafe.Sizeof(int32(10000000)))

	fmt.Println(unsafe.Sizeof(int64(10000000000000)))

	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	fmt.Println(unsafe.Sizeof(string("isson")))

	fmt.Println(unsafe.Sizeof([]string{"isson1", "isson2"}))

}

type person struct {
	Name string
	Age  int
}
