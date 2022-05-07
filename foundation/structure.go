package main

import (
	json2 "encoding/json"
	"fmt"
	"sort"
)

// 自定义类型别名
type myStructure = struct {
	name string
	age  int
}

/**
结构体
type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
    }
*/

type MyStructure0 = struct {
	name string
	age  int
}
type myStructure1 = struct {
	title string
	//继承结构体
	MyStructure0
}

//匿名结构体
type myStructure2 = struct {
	string
	int
	float32
}

//结构体嵌套
type myStructure3 = struct {
	arr myStructure2
}

func testPointer(structure *myStructure) *myStructure {
	return structure
}

func main() {
	var testMyStructure myStructure
	testMyStructure.name = "hyouka"
	testMyStructure.age = 18

	var testMyStructure1 = myStructure{
		name: "hyouka",
		age:  18,
	}
	fmt.Printf("%v\n", testMyStructure)
	fmt.Println(testMyStructure1)
	// 匿名结构体
	var userInfo struct {
		name string
	}
	userInfo.name = "hyouka"
	// 指针类型结构体
	var pointer1 = new(myStructure)
	fmt.Printf("指针类型结构体%T\n", pointer1)    // *struct { name string; age int } 指针类型的结构体
	fmt.Printf("pointer1=%#v\n", pointer1) // &struct { name string; age int }{name:"", age:0}
	//go中支持直接将结构体指针拿来使用
	var pointer2 = new(myStructure)
	pointer2.name = "hyouka"
	fmt.Printf("pointer1=%#v\n", pointer2) // &struct { name string; age int }{name:"hyouka", age:0}
	fmt.Printf("pointer2=%v\n", pointer2.name)
	// fmt.Printf("pointer1=%#v\n", pointer1) 从中可以得出new一个结构体其实就是拿到了他的指针地址
	// 所以&type相当于就是new了一个结构体指针
	var pointer3 = &myStructure{}
	pointer3.name = "hyouka"
	fmt.Printf("pointer3=%T\n", pointer3)  // *struct {name string, age int}
	fmt.Printf("pointer3=%#v\n", pointer3) // &struct {name string, age int}{name: hyouka, age: 0}
	// 也可以通过取地址操作初始化结构体
	var pointer4 = &myStructure{
		name: "hyouka",
		age:  10,
	}
	fmt.Printf("pointer4=%#v\n", pointer4)
	// 通过new操作符不可初始化
	//var pointer5 = new(myStructure{name: "hyouka"})
	// 可以省略结构体的键, 顺序按照结构体从上之下的顺序初始化
	// 通过省略键的方式初始化结构体需要初始化结构体当中所有的字段
	var pointer6 = &myStructure{
		"hyouka",
		18,
	}
	fmt.Printf("pointer6:%#v\n", pointer6) // &struct {name string, age int}{name: "hayouka", age: 18}
	var pointer7 = &myStructure{
		name: "hyouka",
		age:  20,
	}
	fmt.Printf("pointer7.name%p\n", &pointer7.name) //0xc00000c150
	fmt.Printf("pointer7.age%p\n", &pointer7.age)   //0xc00000c160
	//结构体指针地址都是有规律增长
	var pointer8 = struct {
		a, b, c, d int
	}{1, 2, 3, 4}
	fmt.Printf("pointer8.a%p\n", &pointer8.a)
	fmt.Printf("pointer8.b%p\n", &pointer8.b)
	fmt.Printf("pointer8.c%p\n", &pointer8.c)
	fmt.Printf("pointer8.d%p\n", &pointer8.d)
	// 值为指针结构体类型的map
	m := make(map[string]*myStructure) // [map[]]
	stus := []myStructure{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	} // [{pprof.cn, 18}, {}, {}]
	fmt.Printf("stus=%v\n", stus)
	for _, stu := range stus {
		fmt.Printf("当前的值:%p\n", testPointer(&stu))
		// stu为切片的每一项
		// m["pprof.cn"] = {"博客", 28}的指针地址
		m[stu.name] = testPointer(&stu)
	}
	//当前的m=map[pprof.cn:0xc00000c1c8 测试:0xc00000c1c8, 博客:0xc00000c1c8]
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	var pointer9 = &myStructure1{
		title: "hyouka",
		MyStructure0: MyStructure0{
			name: "hyouka1",
			age:  22,
		},
	}
	fmt.Printf("pointer.name=%v\n", pointer9.MyStructure0)
	// 结构体 => json序列化
	json, _ := json2.Marshal(pointer9)
	// {}
	// json序列化时大写开头的键会默认存为key, 小写的则为私有属性
	fmt.Printf("pointer.json => %s\n", json) // {}
	var myStructure4 = struct {
		Name string
		age  int
	}{
		Name: "这是公有的",
		age:  19,
	}
	json1, _ := json2.Marshal(myStructure4)
	// {"Name":"这是公有的"}
	fmt.Printf("pointer.json1 => %s\n", json1)
	// 可以通过``来定义结构体被json转义时的键值定义
	var myStructure5 = struct {
		Name string `json:"name"`
		age  int    `json:"age"`
	}{
		Name: "hyouka",
		age:  20,
	}
	json3, _ := json2.Marshal(myStructure5)
	// {"name":"hyouka"}
	// 同样的age不会被json转义
	fmt.Printf("pointer.json2 => %s\n", json3)
	main1()

	var pointer10 = make(map[int]myStructure)
	pointer10[0] = myStructure{
		name: "1",
		age:  0,
	}
	pointer10[1] = myStructure{
		name: "2",
		age:  0,
	}
	fmt.Printf("pointer10=%v\n", pointer10) //map[0:{}, 1:{}]
	delete(pointer10, 0)
	fmt.Printf("pointer10=%v\n", pointer10) //map[1:{}]
	// 有序输出map
	var sortMap = make(map[int]string, 5)
	sortMap[0] = "0"
	sortMap[1] = "1"
	sortMap[3] = "3"
	sortMap[2] = "2"
	sortMap[4] = "4"
	var sortMapInt []int
	for key := range sortMap {
		sortMapInt = append(sortMapInt, key)
	}
	// 进行排序
	sort.Ints(sortMapInt)
	fmt.Println(sortMapInt)
	for key := range sortMapInt {
		fmt.Println(sortMap[key])
	}
}

// 下面是练习部分

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}

func main1() {
	var ce []student //定义一个切片类型的结构体
	// 初始化切片类型的结构体
	ce = []student{
		{id: 1, name: "阿明", age: 22},
		{id: 2, name: "阿明2", age: 28},
	}
	fmt.Println(ce) //[{id: 1, name: "阿明", age: 22}, {id: 1, name: "阿明", age: 22}]
	demo(ce)
	fmt.Println(ce) //[{id: 1, name: "阿明", age: 22}, {id: 1, name: "阿明", age: 999}]
}
