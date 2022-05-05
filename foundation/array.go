package main

import "fmt"

var arr0 [5]int = [5]int{1, 2, 3} // 长度为5的整数类型
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}           // 不限制长度
var str = [5]string{3: "hello world", 4: "tom"} // 使用初始索引初始化数组
var d = [...]struct {
	name string
	age  int
}{{"hyouka", 16}}

func sumFunc() {
	var sum = 0
	for _, value := range arr2 {
		sum += value
	}
	fmt.Println(sum)
}

// {1, 2, 3, 5, 7, 8} 9
/**
1 other = 9 - 1 = 8;
	2
*/
func targetSum(targetArray [6]int, target int) {
	//for index, value := range targetArray {
	//	other := target - value
	//	for i := index + 1; i < len(targetArray); i++ {
	//		if other == targetArray[i] {
	//			fmt.Println([2]int{i, index})
	//		}
	//	}
	//	println(value)
	//}
	for index, value := range targetArray {
		for i := index + 1; i < len(targetArray); i++ {
			if value+targetArray[i] == target {
				fmt.Println([2]int{index, i})
			}
		}
	}
}

func main() {
	fmt.Println(arr0) // 不够长度时会将不够的部分补齐为0
	fmt.Println(str[3])
	fmt.Println(d[0])
	fmt.Println(d[0].name)
	fmt.Println(d)
	fmt.Println(len(d)) // len cap都会输出数组的长度
	fmt.Println(cap(d))
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	// 使用range 能够获取 slice 或者 map 的下标与其相对应的值
	// index, value := range target
	for key, value := range arr1 {
		fmt.Println(key, value)
	}

	var f = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// _在go中是特殊字符可以用来忽略结果
	for _, value := range f {
		for _, value1 := range value {
			fmt.Println(value1)
		}
	}
	sumFunc()
	targetSum([6]int{1, 2, 3, 5, 7, 8}, 9)
}
