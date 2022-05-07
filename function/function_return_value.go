package main

//命名返回参数可以让函数隐式返回值
func testFunction7(x, y int) (z int) {
	z = x + y
	return
}

func add(x, y int) (z int) {
	defer func() {
		println(z) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
	add(1, 2)
}
