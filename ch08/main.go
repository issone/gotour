package main

import (
	"fmt"
	"time"
)

func main() {

	// 无缓冲 channel，它的容量是 0，不能存储任何数据。所以无缓冲 channel 只起到传输数据的作用，数据并不会在 channel 中做任何停留。这也意味着，无缓冲 channel 的发送和接收操作是同时进行的，它也可以称为同步 channel。
	ch := make(chan string) // 声明一个 channel,  channel 里的数据是 string 类型

	go func() {

		fmt.Println("isson")

		ch <- "goroutine 完成" // 向 chan 发送值，把值放在 chan 中

	}()

	fmt.Println("我是 main goroutine")

	v := <-ch // 获取 chan 中的值, 如果 ch 中没有值，则阻塞等待到 ch 中有值可以接收为止。

	fmt.Println("接收到的chan中的值为：", v)

	// 有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小

	// 一个有缓冲 channel 具备以下特点：
	// 	1. 有缓冲 channel 的内部有一个缓冲队列；
	// 	2. 发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
	// 	3. 接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。

	cacheCh := make(chan int, 5)

	cacheCh <- 2

	cacheCh <- 3

	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))

	close(cacheCh) // 如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。

	fmt.Println("当前值1:", <-cacheCh)
	fmt.Println("当前值2:", <-cacheCh)
	fmt.Println("当前值3:", <-cacheCh)
	x, ok := <-cacheCh // 当从信道中读取数据时，可以有多个返回值，其中第二个可以表示 信道是否被关闭，如果已经被关闭，ok 为 false，若还没被关闭，ok 为true。
	fmt.Println("x:", x)
	fmt.Println("ok:", ok)

	// 单向channel , 限制一个 channel 只可以接收但是不能发送，或者限制一个 channel 只能发送但不能接收，在函数或者方法的参数中，使用单向 channel 的较多，这样可以防止一些操作影响了 channel。
	// onlySend := make(chan<- int)

	// onlyReceive := make(<-chan int)

	// 遍历信道，可以使用 for 搭配 range关键字，在range时，要确保信道是处于关闭状态，否则循环会阻塞。
	pipline := make(chan int, 10)

	go fibonacci(pipline)

	for k := range pipline {
		fmt.Println(k)
	}

	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)
	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()
	go func() {
		secondCh <- downloadFile("secondCh")
	}()
	go func() {
		threeCh <- downloadFile("threeCh")
	}()
	//开始select多路复用，N 个 channel 中，任意一个 channel 有数据产生，select 都可以监听到，然后执行相应的分支，接收数据并处理。	如果一个 select 没有任何 case，那么它会一直等待下去。

	select { //  select 里case不是顺序执行的
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
		// default:		// 如果定义了default，遍历完所有case都没有得到结果，就会进入default分支
		// 	fmt.Println("No data received.")
	}
	// 在 Go 语言中，提倡通过通信来共享内存，而不是通过共享内存来通信，其实就是提倡通过 channel 发送接收消息的方式进行数据传递，而不是通过修改同一个变量。
	// 所以在数据流动、传递的场景中要优先使用 channel，它是并发安全的，性能也不错。

	c1 := make(chan int, 2)

	c1 <- 2
	select { //  select 里的 case 表达式只要求你是对信道的操作即可，不管你是往信道写入数据，还是从信道读出数据
	case c1 <- 4:
		fmt.Println("c1 received: ", <-c1)
		fmt.Println("c1 received: ", <-c1)
	default:
		fmt.Println("channel blocking")
	}

}

func counter(out chan<- int) {
	//函数内容使用变量out，只能进行发送操作
}

func downloadFile(chanName string) string {
	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(mychan)
}
