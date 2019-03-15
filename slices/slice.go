package slices

import "fmt"

var identifier []int

// 切片不需要说明长度 或使用make()函数来创建切片

func test1() {
	var len = 10
	var slice1 []int = make([]int, len)
	slice2 := make([]int, 10)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func test2() {
	s := []int{1, 2, 3}
	fmt.Println(s)
}

func test3() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	s := arr[:]
	fmt.Println(s)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
