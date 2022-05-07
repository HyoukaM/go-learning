package main

import "fmt"

func main() {
	var name = "hyouka"
	// for循环有三种模式
	// for init; condition; post {}
	// for condition{}
	// for{}
	// init 赋值表达式
	// condition 关系表达式或者逻辑表达式
	// post 赋值表达式
	for i := 0; i < len(name); i++ {
		fmt.Println(i)
	}
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}
