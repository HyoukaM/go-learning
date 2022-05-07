package main

import "fmt"

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 循环
	for key := range arr {
		fmt.Println(key)
		// if 表达式(可省略), 判断语句
		if key == 0 {
			fmt.Println("第一次输出")
		} else {
			fmt.Println("第一次以后输出")
		}
		if s := "hello golang"; key < len(s) {
			//切片输出
			fmt.Printf("当前的值%v\n", s[key:key+1])
		}
	}
}
