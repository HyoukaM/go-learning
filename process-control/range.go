package main

import "fmt"

func main() {
	/*
		在go中 range类似于迭代器操作
		对于map,结构体来说返回(键，值)
		对于数组切片来说返回(下标，值)
	*/
	var array = [5]int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("循环数组,下标:%d,值:%d\n", index, value)
	}
	var splice []int
	splice = append(splice, 1)
	var mapArray = make(map[string]string, 2)
	mapArray["张三"] = "李四"
	mapArray["李四"] = "张三"
	for key, value := range mapArray {
		fmt.Printf("循环map,键:%v,值:%v\n", key, value)
	}
	for index, value := range array {
		if value == 1 {
			array[1] = 999
			array[2] = 999
			fmt.Println(array)
		}
		array[index] += 100
	}
	fmt.Println(array)
}
