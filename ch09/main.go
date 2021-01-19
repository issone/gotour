package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源
var sum = 0
var mutex sync.RWMutex

/*
RWMutex，它将程序对资源的访问分为读操作和写操作
为了保证数据的安全，它规定了当有人还在读取数据（即读锁占用）时，不允计有人更新这个数据（即写锁会阻塞）
为了保证程序的效率，多个人（线程）读取数据（拥有读锁）时，互不影响不会造成阻塞，它不会像 Mutex 那样只允许有一个人（线程）读取同一个数据。
读锁：调用 RLock 方法开启锁，调用 RUnlock 释放锁
写锁：调用 Lock 方法开启锁，调用 Unlock 释放锁（和 Mutex类似）
*/

// Mutex 锁的两种定义方法

//  第一种
// var lock *sync.Mutex
// lock = new(sync.Mutex)

//  第二种
// lock := &sync.Mutex{}

func add(i int) {

	mutex.Lock()

	defer mutex.Unlock()

	sum += i

}

func readSum() int {

	//只获取读锁

	mutex.RLock()

	defer mutex.RUnlock()

	b := sum

	return b

}

func run() {

	var wg sync.WaitGroup // sync.WaitGroup 用于最终完成的场景，关键点在于一定要等待所有协程都执行完毕

	//因为要监控110个协程，所以设置计数器为110

	wg.Add(110)

	for i := 0; i < 100; i++ {

		go func() {

			//计数器值减1

			defer wg.Done()

			add(10)

		}()

	}

	for i := 0; i < 10; i++ {

		go func() {

			//计数器值减1

			defer wg.Done()

			fmt.Println("和为:", readSum())

		}()

	}

	//一直等待，只要计数器值为0

	wg.Wait()

}

func doOnce() {
	var once sync.Once // sync.Once 会保证 onceBody 函数只执行一次。sync.Once 适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

//10个人赛跑，1个裁判发号施令

func race() {

	cond := sync.NewCond(&sync.Mutex{}) // 生成一个 *sync.Cond，具有阻塞协程和唤醒协程的功能，所以可以在满足一定条件的情况下唤醒协程

	var wg sync.WaitGroup

	wg.Add(11)

	for i := 0; i < 10; i++ {

		go func(num int) {

			defer wg.Done()

			fmt.Println(num, "号已经就位")

			cond.L.Lock()

			cond.Wait() //等待发令枪响	, 阻塞当前协程，直到被其他协程调用 Broadcast(唤醒所有等待的协程) 或者 Signal (唤醒一个等待时间最长的协程)方法唤醒，使用的时候需要加锁，使用 sync.Cond 中的锁即可，也就是 L 字段

			fmt.Println(num, "号开始跑……")

			cond.L.Unlock()

		}(i)

	}

	//等待所有goroutine都进入wait状态

	time.Sleep(2 * time.Second)

	go func() {

		defer wg.Done()

		fmt.Println("裁判已经就位，准备发令枪")

		fmt.Println("比赛开始，大家准备跑")

		cond.Broadcast() //发令枪响

	}()

	//防止函数提前返回退出

	wg.Wait()

}

func main() {
	run()
	doOnce()
	race()
}
