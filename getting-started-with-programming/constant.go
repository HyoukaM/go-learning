// 包名
package main

import "fmt"

// 常量组
const (
	constantAge    int    = 20
	constantName   string = "hyouka"
	constantHeight int    = 170
	x                     = 12
	y                     // 自动补全 y = 12
	z                     // z = 12
	a, b           = 1, 2
	c, _           // c = 1, _ = 2
)

// 包级变量组
const (
	x1 = 1        // inta = 0
	x2 = iota + 1 // iota 在go中是一个特殊的有名常量默认值为0 1+1 = 2
)

// 主入口函数
func main() {
	// 定义一个变量
	var age = 18
	// 定义了一个
	var name string = "hyouka"
	// 匿名变量自主推导类型
	height := 170
	// 赋值操作
	age = 20

	//age = int(1.3)  类型转换不匹配

	fmt.Println("名字:", name)
	fmt.Println("年龄:", age)
	fmt.Println("身高:", height)

	fmt.Println(constantAge)
	// 局部的常量
	const constantAge1, constantName1, constantHeight1 = 18, "hyouka", 20

	println(x2)

}
