package main

import "fmt"

func main() {

}

// Go 语言有两部分内存空间：栈内存和堆内存。

// 栈内存由编译器自动分配和释放，开发者无法控制。栈内存一般存储函数中的局部变量、参数等，函数创建的时候，这些内存会被自动创建；函数返回的时候，这些内存会被自动释放。

// 堆内存的生命周期比栈内存要长，如果函数返回的值还会在其他地方使用，那么这个值就会被编译器自动分配到堆上。堆内存相比栈内存来说，不能自动被编译器释放，只能通过垃圾回收器才能释放，所以栈内存效率会很高。

// 逃逸到堆内存的变量不能马上被回收，只能通过垃圾回收标记清除，增加了垃圾回收的压力，所以要尽可能地避免逃逸，让变量分配在栈内存上，这样函数返回时就可以回收资源，提升效率

// go build -gcflags="-m -l" ./ch19/main.go  可以使用这个命令进行逃逸分析， 在这一命令中，-m 表示打印出逃逸分析信息，-l 表示禁止内联，可以更好地观察逃逸。

// 几种一定会发生逃逸的情况

// 1. 指针作为函数返回值的时候，一定会发生逃逸

func newString() *string {

	s := new(string)

	*s = "isson"

	return s // 改为返回*s, 可以避免逃逸

}

// 2. 被已经逃逸的指针引用的变量也会发生逃逸

func test2() {
	// isson 这个字符串逃逸到了堆上，这是因为 isson 这个字符串被已经逃逸的指针变量引用，所以它也跟着逃逸了
	fmt.Println("isson")
}

// Go 语言中有 3 个比较特殊的类型，它们是 slice、map 和 chan，被这三种类型引用的指针也会发生逃逸

func test3() {
	m := map[int]*string{}
	s := "isson"
	m[0] = &s // 变量 m 没有逃逸，反而被变量 m 引用的变量 s 逃逸到了堆上。所以被map、slice 和 chan 这三种类型引用的指针一定会发生逃逸的。
}

// 逃逸分析是判断变量是分配在堆上还是栈上的一种方法，在实际的项目中要尽可能避免逃逸，这样就不会被 GC 拖慢速度，从而提升效率
// 从逃逸分析来看，指针虽然可以减少内存的拷贝，但它同样会引起逃逸，所以要根据实际情况选择是否使用指针。

// 优化技巧

// 1. 尽可能避免逃逸，因为栈内存效率更高，还不用 GC。比如小对象的传参，array 要比 slice 效果好
// 2. 如果避免不了逃逸，还是在堆上分配了内存，那么对于频繁的内存申请操作，我们要学会重用内存，比如使用 sync.Pool
// 3. 用合适的算法，达到高性能的目的，比如空间换时间
// 4. 要尽可能避免使用锁、并发加锁的范围要尽可能小、使用 StringBuilder 做 string 和 [ ] byte 之间的转换、defer 嵌套不要太多等

// Go 语言自带的性能剖析的工具 pprof，可以查看 CPU 分析、内存分析、阻塞分析、互斥锁分析
