package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	// test1()
	// test2()
	test3()

}

func test1() {

	a1 := [2]string{"isson", "张三"}

	s1 := a1[0:1]

	s2 := a1[:]

	//打印出s1和s2的Data值，是一样的,也就是这两个切片共用一个数组，所以我们在切片赋值、重新进行切片操作时，使用的还是同一个数组，没有复制原来的元素。这样可以减少内存的占用，提高效率。
	// 多个切片共用一个底层数组虽然可以减少内存占用，但是如果有一个切片修改内部的元素，其他切片也会受影响。所以在切片作为参数在函数间传递的时候要小心，尽可能不要修改原切片内的元素。

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	// 切片的本质是 SliceHeader，又因为函数的参数是值传递，所以传递的是 SliceHeader 的副本，而不是底层数组的副本。
	// 这时候切片的优势就体现出来了，因为 SliceHeader 的副本内存占用非常少，即使是一个非常大的切片（底层数组有很多元素），也顶多占用 24 个字节的内存，这就解决了大数组在传参时内存浪费的问题。
	// SliceHeader 三个字段的类型分别是 uintptr、int 和 int，在 64 位的机器上，这三个字段最多也就是 int64 类型，一个 int64 占 8 个字节，三个 int64 占 24 个字节内存

	// 要获取切片数据结构的三个字段的值，也可以不使用 SliceHeader，而是完全自定义一个结构体，只要字段和 SliceHeader 一样就可以了
	// Data 用来指向存储切片元素的数组。Len 代表切片的长度。Cap 代表切片的容量。
	sh1 := (*slice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.Data, sh1.Len, sh1.Cap)

	// 如果从集合类型的角度考虑，数组、切片和 map 都是集合类型，因为它们都可以存放元素，但是数组和切片的取值和赋值操作要更高效，因为它们是连续的内存操作，通过索引就可以快速地找到元素存储的地址
}

func test2() {
	a1 := [2]string{"isson", "张三"}
	fmt.Printf("函数main数组指针：%p\n", &a1)
	arrayF(a1)
	s1 := a1[0:1]
	fmt.Println("函数main数组的内存地址", (*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)

	// 同一个数组在 main 函数中的指针和在 arrayF 函数中的指针是不一样的，这说明数组在传参的时候被复制了，又产生了一个新数组。
	// 而 slice 切片的底层 Data 是一样的，这说明不管是在 main 函数还是 sliceF 函数中，这两个切片共用的还是同一个底层数组，底层数组并没有被复制
	// 切片的高效还体现在 for range 循环中，因为循环得到的临时变量也是个值拷贝，所以在遍历大的数组时，切片的效率更高。
	// 切片基于指针的封装是它效率高的根本原因，因为可以减少内存的占用，以及减少内存复制时的时间消耗。
}

func test3() {
	s := "isson"

	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

	b := []byte(s)

	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)

	// []byte(s) 和 string(b) 这种强制转换会重新拷贝一份字符串
	s3 := string(b)

	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)

	//  打印出的内存地址都不一样，这说明虽然内容相同，但已经不是同一个字符串了，因为内存地址不同.

	s4 := *(*string)(unsafe.Pointer(&b)) // s4 没有申请新内存（零拷贝），它和变量 b 使用的是同一块内存，因为它们的底层 Data 字段值相同，这样就节约了内存
	fmt.Printf("s4的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)

	// SliceHeader 有 Data、Len、Cap 三个字段，StringHeader 有 Data、Len 两个字段，所以 *SliceHeader 通过 unsafe.Pointer 转为 *StringHeader 的时候没有问题，
	// 因为 *SliceHeader 可以提供 *StringHeader 所需的 Data 和 Len 字段的值。但是反过来却不行了，因为 *StringHeader 缺少 *SliceHeader 所需的 Cap 字段，需要我们自己补上一个默认值。

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)

	// 注意：通过 unsafe.Pointer 把 string 转为 []byte 后，不能对 []byte 修改，比如不可以进行 b1[0]=12 这种操作，会报异常，导致程序崩溃。这是因为在 Go 语言中 string 内存是只读的。

}

func arrayF(a [2]string) {
	fmt.Printf("函数arrayF数组指针：%p\n", &a)
}
func sliceF(s []string) {
	fmt.Printf("函数sliceF 指针：%p , Data：%d\n", s, (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}
