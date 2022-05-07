package main

import "fmt"

func testFunction2(x, y int) (int, int) {
	var temp = x
	x = y
	y = temp
	return x, y
}

// 指针操作
func testFunction3(x, y *int) {
	var temp int //temp为int
	temp = *x    // temp = *&x
	*x = *y      // *&x = *&y
	*y = temp    // *y = *&temp
}

// 多个参数
func testFunction4(n ...int) {

}

// 一个或者多个参数，以此类推
func testFunction5(arg string, n ...int) {

}

func testFunction6(n ...interface{}) {

}

func main() {
	var x, y int = 1, 2
	fmt.Println(testFunction2(x, y))
	fmt.Println(x, y)     // 1, 2 函数进行的是复制操作
	testFunction3(&x, &y) //通过地址操作
	fmt.Println(x, y)     //2,1
}
