package main

import "fmt"

//&（取地址）和*（根据地址取值）
func main() {
	var pointer1 = 10
	var pointer2 = &pointer1 // 取得pointer1的指针值赋值给pointer2
	*pointer2 = 30           // 取到指针值改变成对应的值
	var pointer3 = 30
	pointer3 = pointer1
	pointer3 = 40                      // 赋值操作不会改变pointer1的值
	fmt.Printf("改变指针的值%d\n", pointer1) //30
	fmt.Printf("改变指针的值%d\n", pointer3) //40

	// 空指针
	var str *string
	if str == nil {
		fmt.Println("是空的指针")
	} else {
		fmt.Println("不是空的指针")
	}
	// 下面代码会引起恐慌
	//var a *int
	//*a = 10
	//var b map[string]int
	//b["测试"] = 100
	//fmt.Println(b)
	var a *int
	a = new(int) // 0
	println(*new(int))
	*a = 100
	println(*a)

	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)

	var p int
	println(p)
	var s *int
	s = &p
	*s = 20
	fmt.Println(s, &s, *s)
}
