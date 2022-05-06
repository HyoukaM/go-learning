package main

import "fmt"

func main() {
	// 切片并不是数组或者数组指针
	var slice1 []int
	if slice1 == nil {
		fmt.Println("是空的")
	} else {
		fmt.Println("不是空的")
	}
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice2 = arr[1]     // 指定位置切片
	var slice3 = arr[:]     // 0 => len(arr)-1 长度的所有切片
	var slice4 = arr[1:3]   // 0 => n-1长度的切片 含头不含尾
	var slice5 = arr[1:]    // n => len(arr) - 1 切片
	var slice6 = arr[:4]    // 0 => n - 1 的 切片
	var slice7 = arr[1:3:5] // 0 => n - 1 的 切片
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("slice4: %v\n", slice4)
	fmt.Printf("slice5: %v\n", slice5)
	fmt.Printf("slice6: %v\n", slice6)
	fmt.Printf("slice7: %v\n", slice7)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	var make0 = make([]int, 10)
	// make(type, size, length)
	var make1 = make([]int, 10, 20)
	fmt.Printf("make0: %v\n", make0)
	fmt.Printf("make1: %v\n", make1)
	fmt.Println(cap(make1)) // 总容量为20
	var arr1 = []int{1, 2, 3, 4, 5}
	var slice8 = arr1[1:3] // {2,3}
	slice8[0] += 100       // 2 + 100
	slice8[1] += 100       // 3 + 100
	fmt.Printf("slice8: %v\n", slice8)
	fmt.Printf("arr1: %v\n", arr1) // {1, 102, 103, 4, 5} 会改变原始数组, 操作了底层数组指针
	var arr2 = []int{1, 2, 3}
	var s = &arr2[1] //获取2的指针
	*s += 100
	fmt.Println(*s)
	fmt.Println(arr2[1])
	fmt.Println(arr2) // {1, 102, 3}
	var arr3 = []int{1, 2, 3}
	var arr4 = []int{4, 5, 6}
	var arr5 = append(arr3, arr4...) //向{1,2,3}追加{4,5,6}
	fmt.Printf("arr5: %v\n", arr5)

	s1 := make([]int, 0, 1)
	c := cap(s1)

	for i := 0; i < 50; i++ {
		s1 = append(s1, i)
		// n = 2 n = 2 > c = 1
		if n := cap(s1); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
	fmt.Println(s1)

	var s2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s3 = s2[7:] // {8, 9}
	var s4 = s2[:5] // {1, 2, 3, 4, 5}
	fmt.Printf("%d -> %d\n", s3, s4)
	copy(s4, s3) // 将s3拷贝至s4 {8, 9, 3, 4, 5}
	fmt.Printf("拷贝后的%d->%d\n", s3, s4)
	var str = "hello hyouka"
	var st = str[0:5]
	var name = str[6:]
	fmt.Println("字符串拷贝", st)
	fmt.Println("字符串拷贝", name)
}
