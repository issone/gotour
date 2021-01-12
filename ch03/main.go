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

// switch 后可以不接任何变量、表达式、函数。
// 当不接任何东西时，switch - case 就相当于 if - elseif - else
func test8() {

	score := 30

	switch {
	case score >= 95 && score <= 100:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 60:
		fmt.Println("合格")
	case score >= 0:
		fmt.Println("不合格")
	default:
		fmt.Println("输入有误...")
	}
}

// case 后可以接多个多个条件，多个条件之间是 或 的关系，用逗号相隔。
func test9() {
	month := 2

	switch month {
	case 3, 4, 5:
		fmt.Println("春天")
	case 6, 7, 8:
		fmt.Println("夏天")
	case 9, 10, 11:
		fmt.Println("秋天")
	case 12, 1, 2:
		fmt.Println("冬天")
	default:
		fmt.Println("输入有误...")
	}
}


// 判断一个同学是否有挂科记录的函数,返回值是布尔类型
func getResult(args ...int) bool {
    for _, i := range args {
        if i < 60 {
            return false
        }
    }
    return true
}

// switch 后面可以接一个函数，只要保证 case 后的值类型与函数的返回值 一致即可。
func test10() {
    chinese := 80
    english := 50
    math := 100

    switch getResult(chinese, english, math) {
    // case 后也必须 是布尔类型
    case true:
        fmt.Println("该同学所有成绩都合格")
    case false:
        fmt.Println("该同学有挂科记录")
}