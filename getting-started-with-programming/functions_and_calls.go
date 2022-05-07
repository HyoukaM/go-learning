package main

import "fmt"

// 局部函数(接受的参数)(返回参数)
func testFunction(a int, b int) (s int, d int) {
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return s, d
}

// 包名函数可在其他文件中调用
func VersionString() string {
	return "go1.0"
}

func main() {
	s, d := testFunction(10, 20)
	fmt.Println(s, d)
	println(VersionString())
	// 匿名函数
	x, y := func() (int, int) {
		println("这是立即执行函数")
		return 3, 4
	}() // 一对小括号表示立即调用此函数。不需传递实参。
	println(x, y)
}
