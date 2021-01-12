package main

import "fmt"

// Age 定义一个uint的Age类型
type Age uint

func (age Age) String() { // 方法 String() 就是类型 Age 的方法，类型 Age 是方法 String() 的接收者

	fmt.Println("the age is", age)

}

// Modify 修改年龄 ,方法的接收者除了可以是值类型（比如上一小节的示例），也可以是指针类型
func (age *Age) Modify() {

	*age = Age(30)

}

// 结构体
type person struct {
	name string

	age uint
}

type person2 struct {
	name string
	age  uint
	addr address // 结构体的字段可以是任意类型，也包括自定义的结构体类型
}
type address struct {
	province string
	city     string
}

// type Stringer interface {	// Stringer 是 Go SDK 的一个接口，属于 fmt 包。
//     String() string
// }

// person 就实现了 Stringer 接口, 注意：如果一个接口有多个方法，那么需要实现接口的每个方法才算是实现了这个接口
func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// New 工厂函数，返回一个error接口，其实具体实现是*errorString
func New(text string) error {

	return &errorString{text}

}

//结构体，内部一个字段s，存储错误信息

type errorString struct {
	s string
}

//用于实现error接口

func (e *errorString) Error() string {

	return e.s

}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter 是Reader和Writer的组合
// 类型组合后，外部类型不仅可以使用内部类型的字段，也可以使用内部类型的方法，就像使用自己的方法一样。如果外部类型定义了和内部类型同样的方法，那么外部类型的会覆盖内部类型，这就是方法的覆写
type ReadWriter interface {
	Reader

	Writer
}

func main() {
	// 25 也是 unit 类型，unit 类型等价于定义的 Age 类型，所以 25 可以强制转换为 Age 类型。
	age := Age(25)
	age.String()

	//在调用方法的时候，传递的接收者本质上都是副本，只不过一个是这个值副本，一是指向这个值指针的副本。指针具有指向原有值的特性，所以修改了指针指向的值，也就修改了原有的值。我们可以简单地理解为值接收者使用的是值的副本来调用方法，而指针接收者使用实际的值来调用方法。
	age.Modify()
	age.String()
	// 也可以使用指针变量调用，Go 语言编译器帮我们自动做的事情
	// 如果使用一个值类型变量调用指针类型接收者的方法，Go 语言编译器会自动帮我们取指针调用，以满足指针接收者的要求。
	// 同样的原理，如果使用一个指针类型变量调用值类型接收者的方法，Go 语言编译器会自动帮我们解引用调用，以满足值类型接收者的要求。
	(&age).Modify()

	//方法赋值给变量，方法表达式
	sm := Age.String
	//通过变量，要传一个接收者进行调用也就是age
	sm(age)

	// 结构体声明使用
	//var p person	// 没有对变量 p 初始化，所以默认会使用结构体里字段的零值。
	p := person{"xxx", 18} // 类似python里的按位置传参
	fmt.Println(p.name, p.age)
	p1 := person{age: 20, name: "xx"} // 类似python里的按关键字传参
	fmt.Println(p1.name, p1.age)

	p2 := person{age: 30} // 只初始化字段 age，字段 name 使用默认的零值
	fmt.Println(p2.name, p2.age)

	p3 := person2{

		age: 30,

		name: "abc",

		addr: address{

			province: "北京",

			city: "北京",
		},
	}

	fmt.Println(p3.addr.province)

	printString(p)  //  person 实现了 Stringer 接口，所以变量 p 可以作为函数 printString 的参数
	printString(&p) // 把变量 p 的指针作为实参传给 printString 函数也是可以的
	// 如果实现接口时，以指针*p作为接受者，则实参只能传递指针类型
	// 以值类型接收者实现接口的时候，不管是类型本身，还是该类型的指针类型，都实现了该接口。
	// 以指针类型接收者实现接口的时候，只有对应的指针类型才被认为实现了该接口

	var s fmt.Stringer

	a, ok := s.(address)
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s不是一个address")
	}

}
