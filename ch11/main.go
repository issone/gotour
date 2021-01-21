package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// selectTimeoutTest()
	// pipelineTest()
	// mpTest()
	futereTest()
}

// for select 循环模式

func fsTest() {
	//  for + select 无限循环，这种模式会一直执行 default 语句中的任务，直到 done 这个 channel 被关闭为止
	done := make(chan bool)

	for {

		select {

		case <-done:

			return

		default:

			//执行具体的任务

		}

	}

}

func frsTest() {
	//  for range select 有限循环,
	done := make(chan bool)
	resultCh := make(chan int)

	for _, s := range []int{} {

		select {

		case <-done: // 用于退出当前的 for 循环

			return

		case resultCh <- s: // 用于接收 for range 循环的值，这些值通过 resultCh 可以传送给其他的调用者

		}

	}

}

// select timeout 模式,核心在于通过 time.After 函数设置一个超时时间，防止因为异常造成 select 语句的无限等待。如果可以使用 Context 的 WithTimeout 函数超时取消，要优先使用。
func selectTimeoutTest() {
	result := make(chan string)
	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}
}

// Pipeline 模式, Pipeline 流水线模式中的工序是相互依赖的

func pipelineTest() {
	coms := buy(10)       //采购10套配件
	phones := build(coms) //组装10部手机
	packs := pack(phones) //打包它们以便售卖
	//输出测试，看看效果
	for p := range packs {
		fmt.Println(p)
	}
}

//工序1采购

func buy(n int) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for i := 1; i <= n; i++ {

			out <- fmt.Sprint("配件", i)

		}

	}()

	return out

}

//工序2组装

func build(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "组装(" + c + ")"

		}

	}()

	return out

}

//工序3打包

func pack(in <-chan string) <-chan string {

	out := make(chan string)

	go func() {

		defer close(out)

		for c := range in {

			out <- "打包(" + c + ")"

		}

	}()

	return out

}

// 扇入扇出模式, 是为了扩展Pipeline中某个环节的效率

func mpTest() {
	coms := buy(100) //采购100套配件
	//三班人同时组装100部手机
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)
	//汇聚三个channel成一个
	phones := merge(phones1, phones2, phones3)
	packs := pack(phones) //打包它们以便售卖
	//输出测试，看看效果
	for p := range packs {
		fmt.Println(p)
	}
}

//扇入函数（组件），把多个chanel中的数据发送到一个channel中

func merge(ins ...<-chan string) <-chan string {

	var wg sync.WaitGroup

	out := make(chan string)

	//把一个channel中的数据发送到out中

	p := func(in <-chan string) {

		defer wg.Done()

		for c := range in {

			out <- c

		}

	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处理多个channel中的数据

	for _, cs := range ins {

		go p(cs)

	}

	//等待所有输入的数据ins处理完，再关闭输出out

	go func() {

		wg.Wait()

		close(out)

	}()

	return out

}

// Futures 模式， 可以理解为未来模式，主协程不用等待子协程返回的结果，可以先去做其他事情，等未来需要子协程结果的时候再来取，如果子协程还没有返回结果，就一直等待

func futereTest() {
	vegetablesCh := washVegetables() //洗菜
	waterCh := boilWater()           //烧水
	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)

	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了，可以做火锅了:", vegetables, water)
}

//洗菜

func washVegetables() <-chan string {

	vegetables := make(chan string)

	go func() {

		time.Sleep(5 * time.Second)

		vegetables <- "洗好的菜"

	}()

	return vegetables

}

//烧水

func boilWater() <-chan string {

	water := make(chan string)

	go func() {

		time.Sleep(5 * time.Second)

		water <- "烧开的水"

	}()

	return water

}
