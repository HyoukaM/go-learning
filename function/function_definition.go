package main

// func functionName(接受参数定义)(返回类型)
// 如果接受参数当中相邻的两个类型一致可以写作一个, 返回值为多个时需要通过()包裹
// go函数左大括号不能另起一行
func testFunction(x, y int, s string) (int, string) {
	return x + y, s
}

type myFunctionType func(x, y int, s string) string

// 如果参数是 函数类型的复杂类型建议将复杂签名定义为函数类型
/**
@link myFunctionType
*/
func testFunction1(fc myFunctionType, x, y int, s string) string {
	return fc(x, y, s)
}

func main() {
	testFunction(10, 11, "hello golang")
}
