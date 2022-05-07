package main

import "fmt"

func a() func() int {
	var i = 0
	var b = func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func a1() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

// 数字阶乘
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

//斐波那契数列
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	var c = a()
	c()
	c()
	c()
	a()

	a1()
	a1()() // 两个内存地址是一样的
	var i = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
	var ii = 10
	for i := 0; i < ii; i++ {
		fmt.Printf("%d\n", fibonaci(i))
	}
}
