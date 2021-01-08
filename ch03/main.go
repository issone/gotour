package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
}

func test1() {
	i := 10
	if i > 10 {
		fmt.Println("i>10")
	} else {
		fmt.Println("i<=10")
	}

}

func test2() {

	if i := 6; i > 10 {

		fmt.Println("i>10")

	} else if i > 5 && i <= 10 {

		fmt.Println("5<i<=10")

	} else {

		fmt.Println("i<=5")

	}

}

func test3() {
	switch i := 6; {

	case i > 10: // case 后，自带break

		fmt.Println("i>10")

	case i > 5 && i <= 10:

		fmt.Println("5<i<=10")

	default:

		fmt.Println("i<=5")

	}

}

func test4() {
	switch j := 1; j { //当 switch 之后有表达式时，case 后的值就要和这个表达式的结果类型相同

	case 1:

		fallthrough // 不break，继续往下执行判断

	case 2:

		fmt.Println("1")

	default:

		fmt.Println("没有匹配")

	}

}

func test5() {

	sum := 0

	for i := 1; i <= 100; i++ {

		sum += i

	}

	fmt.Println("the sum is", sum)

}

func test6() {

	sum := 0

	i := 1

	for i <= 100 { // 等价于 while i <= 100, go中没有while

		sum += i

		i++

	}

	fmt.Println("the sum is", sum)

}

func test7() {
	sum := 0
	i := 1
	for { // 无限循环
		sum += i
		i++
		if i > 100 {
			break
		}
	}
	fmt.Println("the sum is", sum)
}
