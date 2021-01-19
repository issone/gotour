package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	selectChannel()
	withCancelTest()
	withDeadlineTest()
	withTimeoutTest()
	withValueTest()
}

func watchDog(stopCh chan bool, name string) {

	//开启for select循环，一直后台监控

	for {

		select {

		case <-stopCh:

			fmt.Println(name, "停止指令已收到，马上停止")

			return

		default:

			fmt.Println(name, "正在监控……")

		}

		time.Sleep(1 * time.Second)

	}

}

// 使用select + channel, 让协程自行退出
func selectChannel() {
	var wg sync.WaitGroup

	wg.Add(1)

	stopCh := make(chan bool) //用来停止监控狗

	go func() {

		defer wg.Done()

		watchDog(stopCh, "【监控狗1】")

	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒

	stopCh <- true //发停止指令

	wg.Wait()
}

// 使用 Context 取消多个协程

// Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作
// 根Context有两个，一个是Background，主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context，它不能被取消。一个是TODO，如果我们不知道该使用什么Context的时候，可以使用这个

func monitor(ctx context.Context, number int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("监控器%v，接收到通道值为：%v，监控结束。\n", number, v)
			return
		default:
			// 获取 item 的值
			value := ctx.Value("item")
			fmt.Printf("监控器%v，正在监控 %v \n", number, value)
			time.Sleep(2 * time.Second)
		}
	}
}

// WithCancel(parent Context)：生成一个可取消的 Context。

// WithDeadline(parent Context, d time.Time)：生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。

// WithTimeout(parent Context, timeout time.Duration)：生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消

// WithValue(parent Context, key, val interface{})：生成一个可携带 key-value 键值对的 Context。

// 这四个函数有一个共同的特点，就是第一个参数，都是接收一个 父context。

// 通过一次继承，就多实现了一个功能，比如使用 WithCancel 函数传入 根context ，就创建出了一个子 context，该子context 相比 父context，就多了一个 cancel context 的功能。

// 如果此时，我们再以上面的子context（context01）做为父context，并将它做为第一个参数传入WithDeadline函数，获得的子子context（context02），相比子context（context01）而言，又多出了一个超过 deadline 时间后，自动 cancel context 的功能。

func withCancelTest() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(1 * time.Second)
	// 关闭所有 goroutine
	cancel()

	// 等待5s，若此时屏幕没有输出 <正在监控中> 就说明所有的goroutine都已经关闭
	time.Sleep(5 * time.Second)

	fmt.Println("主程序退出！！")
}

func withDeadlineTest() {
	ctx01, cancel := context.WithCancel(context.Background())

	// WithDeadline 传入的第二个参数是 time.Time 类型，它是一个绝对的时间，意思是在什么时间点超时取消
	ctx02, cancel := context.WithDeadline(ctx01, time.Now().Add(1*time.Second))

	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor(ctx02, i)
	}

	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx02.Err())
	}

	fmt.Println("主程序退出！！")
}

func withTimeoutTest() {
	ctx01, cancel := context.WithCancel(context.Background())

	// WithTimeout 传入的第二个参数是 time.Duration 类型，它是一个相对的时间，意思是多长时间后超时取消
	ctx02, cancel := context.WithTimeout(ctx01, 1*time.Second)

	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor(ctx02, i)
	}

	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx02.Err())
	}

	fmt.Println("主程序退出！！")
}

func withValueTest() {
	ctx01, cancel := context.WithCancel(context.Background())
	ctx02, cancel := context.WithTimeout(ctx01, 1*time.Second)
	// 以 ctx02 为父 context，再创建一个能携带 value 的ctx03，由于他的父context 是 ctx02，所以 ctx03 也具备超时自动取消的功能
	ctx03 := context.WithValue(ctx02, "item", "CPU")
	defer cancel()

	for i := 1; i <= 5; i++ {
		go monitor(ctx03, i)
	}

	time.Sleep(5 * time.Second)
	if ctx02.Err() != nil {
		fmt.Println("监控器取消的原因: ", ctx02.Err())
	}

	fmt.Println("主程序退出！！")
}

// 要更好地使用 Context，有一些使用原则需要尽可能地遵守。
// 通常 Context 都是做为函数的第一个参数进行传递（规范性做法），并且变量名建议统一叫 ctx
// Context 是线程安全的，可以放心地在多个 goroutine 中使用。
// 当你把 Context 传递给多个 goroutine 使用时，只要执行一次 cancel 操作，所有的 goroutine 就可以收到 取消的信号
// 不要把原本可以由函数参数来传递的变量，交给 Context 的 Value 来传递。
// 当一个函数需要接收一个 Context 时，但是此时你还不知道要传递什么 Context 时，可以先用 context.TODO 来代替，而不要选择传递一个 nil。
// 当一个 Context 被 cancel 时，继承自该 Context 的所有 子 Context 都会被 cancel。
