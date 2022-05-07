package main

import (
	"fmt"
	"io"
	"os"
)

/**
	defer特性
	1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
*/
/**
	defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
*/

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}
func delayedFunction() {
	i := [5]struct{ name string }{}
	for index := range i {
		fmt.Println(index) // 0 1 2 3 4
	}
	for index := range i {
		defer fmt.Println(index) // 4 3 2 1 0
	}
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}
}

func do() error {
	f, err := os.Open("function_return_value.go")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}

func main() {
	err := do()
	if err != nil {
		return
	}
}
