package main

import "fmt"

func main() {

	var i = 0

	for i := 0; i < 10; i++ {
		if i == 5 {
			// 跳出当前循环
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			// 跳出此次循环
			continue
		}
		fmt.Println(i)
	}

Exit:
	i++
	if i < 5 {
		goto Exit
	}
}
