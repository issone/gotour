package main

import (
	"errors"
	"fmt"
)

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能为负数") //  errors.New 这个工厂函数可以生成我们想要的错误信息，但是携带的信息只有字符串
	} else {
		return a + b, nil
	}
}

// 自定义error
type commonError struct {
	errorCode int //错误码

	errorMsg string //错误信息

}

func (ce *commonError) Error() string { // 实现Error接口

	return ce.errorMsg

}

func add1(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"} //  使用自定义的 error，可以携带更多的信息
	} else {
		return a + b, nil
	}
}

func main() {

	name1 := "go"
	defer fmt.Println(name) // 输出: go	, 使用 defer 只是延时调用函数，此时传递给函数里的变量，不应该受到后续程序的影响

	name1 = "python"
	fmt.Println(name)

	name1 = "java"
	fmt.Println(name1) // 多个defer 反序调用, 类似栈一样，后进先出

	myname := myfunc()                       // name 此时还是全局变量，值还是go
	fmt.Printf("main 函数里的name: %s\n", name)  // 在 defer 里改变了这个全局变量，此时name的值已经变成了 python , 如果在main中重新赋值了name，则name是局部变量的值
	fmt.Println("main 函数里的myname: ", myname) //  defer 是return 后才调用的。所以在执行 defer 前，myname 已经被赋值成 go 了

	// defer 函数一般用来释放资源，可以保证文件关闭后一定会被执行，不管你自定义的函数出现异常还是错误,

	testError()
	testPanic()
}

func testError() {
	sum, err := add1(-1, 2)

	if cm, ok := err.(*commonError); ok { // 断言 err是否为*commonError类型
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)

	} else {

		fmt.Println(sum)

	}

	// error嵌套：	不想原 error 的情况下，又想添加一些额外信息返回新的 error
	e := errors.New("原始错误e")

	w := fmt.Errorf("Wrap了一个错误:%w", e)

	fmt.Println(w)

	// errors.Unwrap 函数, 获取被嵌套的原始错误 e
	fmt.Println(errors.Unwrap(w))

	// errors.Is 函数: 用来判断两个 error 是否是同一个。  如果 err 和 target 是同一个，那么返回 true；如果 err 是一个 wrapping error，target 也包含在这个嵌套 error 链中的话，也返回 true。

	fmt.Println(errors.Is(w, e))

	// errors.As 函数: error嵌套后不能直接断言，可以使用此方法
	var cm *commonError

	if errors.As(err, &cm) {

		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)

	} else {

		fmt.Println(sum)

	}

	// 在 Go 语言提供的 Error Wrapping 能力下，我们写的代码要尽可能地使用 Is、As 这些函数做判断和转换。
}

// panic 异常是一种非常严重的情况，会让程序中断运行，使程序崩溃，所以如果是不影响程序运行的错误，不要使用 panic，使用普通错误 error 即可。

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
	//省略其他代码
}

func testPanic() {
	// connectMySQL("", "abc", "123")

	//  在 Go 语言中，可以通过内置的 recover 函数恢复 panic 异常。因为在程序 panic 异常崩溃的时候，只有被 defer 修饰的函数才能被执行，所以 recover 函数要结合 defer 关键字使用才能生效
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("", "root", "123456")
}
