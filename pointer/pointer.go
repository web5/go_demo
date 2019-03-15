package pointer

import "fmt"

var ip *int     //  指针变量
var fp *float32 //  指针变量

func test1() {
	var a int = 10
	fmt.Printf("变量的地址: %x\n", &a)
}

func test2() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

// 指针数组
const MAX int = 3

func pointerArr() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d \n", i, *ptr[i])
	}

}

func pointerParams() {
	var a int = 100
	var b int = 200
	fmt.Printf("a : %d , b: %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a : %d , b: %d\n", a, b)
}

// 指针参数
func swap(x *int, y *int) {
	*x, *y = *y, *x
}
