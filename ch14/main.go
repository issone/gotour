package main

import "fmt"

func main() {

	// 变量的声明用var
	var s string
	var sp *string

	fmt.Println("s", &s)  // 并没有对s初始化，但是可以通过 &s 获取它的内存地址
	fmt.Println("sp", sp) // 指针类型在声明的时候，Go 语言并没有自动分配内存, 其值为nil
	// 对于值类型来说，即使只声明一个变量，没有对其初始化，该变量也会有分配好的内存

	// 变量的初始化（赋值），有3种方法
	// 1. 声明时直接初始化，比如 var s string = "123"
	// 2. 声明后再进行初始化，比如 s="123"（假设已经声明变量 s）。
	// 3. 使用 := 简单声明，比如 s:="123"。

	s = "张三"
	fmt.Println(s)

	// 如果要对一个变量赋值，这个变量必须有对应的分配好的内存，这样才可以对这块内存操作，完成赋值的目的

	// *sp = "isson"	// sp是指针类型，声明后没有分配内存，赋值操作会失败， map 和 chan 也一样，因为它们本质上也是指针类型。
	// fmt.Println(*sp)

	// new函数

	sp = new(string) //关键点， new的作用就是根据传入的类型申请一块内存，然后返回指向这块内存的指针，指针指向的数据就是该类型的零值
	fmt.Println("*sp1", *sp)
	*sp = "isson"
	fmt.Println("*sp2", *sp)

	//字面量初始化
	p := person{name: "张三", age: 18}
	fmt.Println("p", p)

	pp := newPerson("isson", 20)
	fmt.Println("name为", pp.name, ",age为", pp.age)

	m := make(map[string]int, 10)
	fmt.Println("m", m)

	// new 函数只用于分配内存，并且把内存清零，也就是返回一个指向对应类型零值的指针。new 函数一般用于需要显式地返回指针的情况，不是太常用。

	// make 函数只用于 slice、chan 和 map 这三种内置类型的创建和初始化，因为这三种类型的结构比较复杂，比如 slice 要提前初始化好内部元素的类型，slice 的长度和容量等，这样才可以更好地使用它们。
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	// 指针变量初始化
	p := new(person)
	p.name = name
	p.age = age
	return p
}
