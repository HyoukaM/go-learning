package main

import "fmt"

func main() {
	var (
		a, b float32 = 12.0, 3.14
		c, d int16   = 15, -6
		e    uint8   = 7
	)
	_ = 12 + 'A'
	x := 12 - a // 12将被当做a的类型（float32）使用。
	fmt.Println(a, b, c, d, e)
	println(x)
}
