package main

import "fmt"

func main() {

	var interfaceType interface{}

	switch interfaceType.(type) {
	case int:
		fmt.Println("这是int类型")
	case float32:
		fmt.Println("这是float32类型")
	default:
		fmt.Println("不知道当前类型")
	}

}
