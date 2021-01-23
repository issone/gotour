package main

import "fmt"

// 在 Go 语言中，函数的参数传递只有值传递，而且传递的实参都是原始数据的一份拷贝。
// 如果拷贝的内容是值类型的，那么在函数中就无法修改原始数据；如果拷贝的内容是指针（或者可以理解为引用类型 map、chan 等），那么就可以在函数中修改原始数据。

// 值类型在参数传递中无法修改, struct 外，还有浮点型、整型、字符串、布尔、数组，这些都是值类型
// 指针类型的变量保存的值就是数据对应的内存地址，所以在函数参数传递是传值的原则下，拷贝的值也是内存地址

func main() {

	p := person{name: "张三", age: 18}

	fmt.Printf("main函数：p的内存地址为%p\n", &p)

	modifyPerson(p)

	fmt.Println("person name:", p.name, ",age:", p.age)

	m := make(map[string]int)

	m["issone"] = 18

	fmt.Println("issone的年龄为", m["issone"])

	modifyMap(m)

	fmt.Println("issone的年龄为", m["issone"])

}

type person struct {
	name string
	age  int
}

func modifyPerson(p person) {

	fmt.Printf("modifyPerson函数：p的内存地址为%p\n", &p)

	p.name = "李四"

	p.age = 20

}

// 严格来说，Go 语言没有引用类型，但是我们可以把 map、chan 称为引用类型，这样便于理解。除了 map、chan 之外，Go 语言中的函数、接口、slice 切片都可以称为引用类型。指针类型也可以理解为是一种引用类型
func modifyMap(p map[string]int) {

	p["issone"] = 20

}
