package main

import "fmt"

func main() {
	// map[keyType]valueType
	// map和指针定义一样
	var s map[string]int
	s = make(map[string]int, 10)
	s["这是键"] = 100
	fmt.Println(s)        // map[这是键: 100]
	fmt.Println(s["这是键"]) // 100
	// map可以初始定值
	var useInfo = map[string]string{
		"name": "hyouka",
		"age":  "20",
	}
	fmt.Println(useInfo)
	// 判断map中是否存在这个值
	_, key := useInfo["name"]
	if key {
		println("存在")
	} else {
		println("不存在")
	}
	// 同样使用 for range来遍历map
	for key, value := range useInfo {
		println(key)
		println(value)
	}
	for key := range useInfo {
		println(key)
	}
	delete(useInfo, "name")
	fmt.Println(useInfo)
	// 元素为map的切片
	var mySliceMap = make([]map[string]string, 2) // {map[] map[]}
	fmt.Println(mySliceMap)
	for key := range mySliceMap {
		fmt.Println(key)
	}
	mySliceMap[0] = make(map[string]string, 2)
	mySliceMap[0]["张三"] = "李四"
	fmt.Println(mySliceMap)
}
