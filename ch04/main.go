package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
}

// 数组的定义及循环
func test1() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("arr", arr)
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("array", array)

	array1 := [...]string{"a", "b", "c", "d", "e"} // 省略数组长度，根据{}中的数量推导
	fmt.Println("array1", array1)

	array2 := [5]string{1: "b", 3: "d"} // 只针对特定索引元素初始化,没有初始化的索引，其默认值都是数组类型的零值，这里是 string 类型的零值
	fmt.Println("array2", array2)

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}

	for i, v := range array { // range 表达式返回两个结果,第一个是数组的索引；第二个是数组的值。 如果返回的值用不到，可以使用 _ 下划线丢弃
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}
}

//  切片
func test2() {
	// 基于数组生成切片
	array := [5]string{"a", "b", "c", "d", "e"}

	slice := array[2:5] // 切片是一个具备三个字段的数据结构，分别是指向数组的指针 data，长度 len 和容量 cap ,这里data指针执行array索引为2的位置， len为3, cap为3

	fmt.Println(slice)

	// 切片修改, 切片的底层是数组 ,修改切片的值，会修改原数组
	slice[1] = "f"
	fmt.Println(array)

	// 切片声明
	slice1 := make([]string, 4) // 声明了一个元素类型为 string 的切片，长度是 4
	slice1 = append(slice1, "A1")
	slice1 = append(slice1, "A2")
	slice1 = append(slice1, "A3")
	fmt.Println(slice1)

	// 切片的容量不能比切片的长度小。 ，Go 语言在内存上划分了一块容量为 8 的内容空间（容量为 8），
	// 但是只有 4 个内存空间才有元素（长度为 4），其他的内存空间处于空闲状态，
	// 当通过 append 函数往切片中追加元素的时候，会追加到空闲的内存上，当切片的长度要超过容量的时候，会进行扩容
	// slice1 := make([]string, 4, 8)

	//追加一个元素
	slice2 := append(slice1, "f")
	fmt.Println(slice2)
	//多加多个元素
	slice2 = append(slice1, "f", "g")
	fmt.Println(slice2)
	//追加另一个切片
	slice2 = append(slice1, slice...) // append 会自动处理切片容量不足需要扩容的问题。
	fmt.Println(slice2)

	//小技巧：在创建新切片的时候，最好要让新切片的长度和容量一样，这样在追加操作的时候就会生成新的底层数组，从而和原有数组分离，就不会因为共用底层数组导致修改内容的时候影响多个切片。

	//切片不仅可以通过 make 函数声明，也可以通过字面量的方式声明和初始化，如下所示：

	slice3 := []string{} // 声明并初始化
	fmt.Println(slice3)

	fmt.Println(len(slice3), cap(slice3))
	slice3 = append(slice3, "1", "2")
	fmt.Println("slice3", slice3)

	var myarr []string //由于 切片是引用类型，所以你不对它进行赋值的话，它的零值（默认值）是 nil
	fmt.Println(myarr == nil)
	myarr = append(myarr, "123")
	fmt.Println(myarr)

}

// MAP

func test3() {
	nameAgeMap := make(map[string]int) // 声明并初始化字典的方法
	nameAgeMap["a"] = 20
	fmt.Println("nameAgeMap", nameAgeMap)
	// 除了可以通过 make 函数创建 map 外，还可以通过字面量的方式
	nameAgeMap2 := map[string]int{}
	fmt.Println(nameAgeMap2)
	nameAgeMap2["aaa"] = 20
	//获取指定 Key 对应的 Value
	age := nameAgeMap2["aaa"]
	fmt.Println("age", age)

	//	map 可以获取不存在的 K-V 键值对，如果 Key 不存在，返回的 Value 是该类型的零值，比如 int 的零值就是 0。所以很多时候，我们需要先判断 map 中的 Key 是否存在。
	//	map 的 [] 操作符可以返回两个值：
	//	第一个值是对应的 Value；
	//	第二个值标记该 Key 是否存在，如果存在，它的值为 true
	age, ok := nameAgeMap["aaa"]
	fmt.Println("ok", ok)
	if ok {
		fmt.Println(age)
	}

	// 要删除 map 中的键值对
	delete(nameAgeMap, "a")
	fmt.Println("nameAgeMap", nameAgeMap)

	// 遍历 Map
	nameAgeMap["a"] = 1
	nameAgeMap["b"] = 2
	nameAgeMap["c"] = 3
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}

	for k := range nameAgeMap { // 使用一个返回值的时候，这个返回值默认是 map 的 Key。
		fmt.Println("k is", k)
	}

	// map 的遍历是无序的，也就是说你每次遍历，键值对的顺序可能会不一样。如果想按顺序遍历，可以先获取所有的 Key，并对 Key 排序，然后根据排序好的 Key 获取对应的 Value.

	fmt.Println(len(nameAgeMap)) // map 是没有容量的，它只有长度，也就是 map 的大小（键值对的个数）

	// 如果我们用var 先声明一个字典，需要注意以下情况

	// 声明一个名为 score 的字典
	var scores map[string]int
	// scores["chinese"] = 80	// 会失败，因为未初始化的 score 的零值为nil，无法直接进行赋值

	if scores == nil {
		// 需要使用 make 函数先对其初始化
		scores = make(map[string]int)
	}

	// 经过初始化后，就可以直接赋值
	scores["chinese"] = 90
	fmt.Println(scores)

}
