package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// 在 Go 反射中，标准库为我们提供了两种类型 reflect.Value 和 reflect.Type 来分别表示变量的值和类型，并且提供了两个函数 reflect.ValueOf 和 reflect.TypeOf 分别获取任意对象的 reflect.Value 和 reflect.Type
// interface{} 是空接口，可以表示任何类型，也就是说你可以把任何类型转换为空接口，它通常用于反射、类型断言，以减少重复代码，简化编程
func main() {

	i := 3

	//int to reflect.Value
	iv := reflect.ValueOf(i) //  reflect.ValueOf , 获得变量对应的 reflect.Value  ,  reflect.Value 表示的是变量的值

	it := reflect.TypeOf(i) //  reflect.TypeOf , 获得变量对应的 reflect.Type  ,  reflect.Value 表示的是变量的类型

	fmt.Println(iv, it) //3 int

	// 获取原始类型, 通过 reflect.ValueOf 函数把任意类型的对象转为一个 reflect.Value，而如果想逆向转回来也可以，reflect.Value 为我们提供了 Inteface 方法
	i1 := iv.Interface().(int) //reflect.Value to int

	fmt.Println(i1)

	// 要修改一个变量的值，有几个关键点：传递指针（可寻址），通过 Elem 方法获取指向的值，才可以保证值可以被修改
	ipv := reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)

	// 	那么如何修改 struct 结构体字段的值呢？参考变量的修改方式，可总结出以下步骤：

	// 传递一个 struct 结构体的指针，获取对应的 reflect.Value；

	// 通过 Elem 方法获取指针指向的值；

	// 通过 Field 方法获取要修改的字段,该字段需要是可导出的，而不是私有的，也就是该字段的首字母为大写；

	// 通过 Set 系列方法修改成对应的值。
	p := person{Name: "isson", Age: 20}
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)

	ppv := reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("张三")
	fmt.Println(p)

	// 遍历结构体的字段和方法

	//遍历person的字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段：", pt.Field(i).Name)
	}
	//遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法：", pt.Method(i).Name)
	}

	// 可以通过 FieldByName 方法获取指定的字段，也可以通过 MethodByName 方法获取指定的方法，这在需要获取某个特定的字段或者方法时非常高效，而不是使用遍历

	// 尽可能通过类型断言的方式判断是否实现了某接口，而不是通过反射
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：", pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：", pt.Implements(writerType))

	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(reflect.TypeOf(jsonB))
		fmt.Println(string(jsonB))

	}
	//json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)

	//遍历person字段中key为json的tag
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n", sf.Name, sf.Tag.Get("json")) // 通过不同的 Key，使用 Get 方法就可以获得自定义的不同的 tag。
		fmt.Printf("字段%s上,bson tag为%s\n", sf.Name, sf.Tag.Get("bson"))
	}

	//自己实现的struct to json
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json") //获取json tag
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		//获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String()) //打印json字符串

	// 总结
	// 1. 任何类型的变量都可以转换为空接口 intferface{}， reflect.ValueOf 和 reflect.TypeOf 的参数就是 interface{}，表示可以把任何类型的变量转换为反射对象
	// 2.reflect.Value 结构体的 Interface 方法返回的值也是 interface{}，表示可以把反射对象还原为对应的类型变量。
	// 3.要修改反射的对象，该值必须可设置，也就是可寻址

	//反射调用person的Print方法
	mPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args) // Call参数类型 []reflect.Value
}

type person struct {
	Name string `json:"name" bson:"b_name"` // 多个 tag，要使用空格分隔
	Age  int    `json:"age" bson:"b_age"`
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}

func (p person) Print(prefix string) {
	fmt.Printf("%s:Name is %s,Age is %d\n", prefix, p.Name, p.Age)
}
