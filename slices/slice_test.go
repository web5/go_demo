package slices

import (
	"fmt"
	"testing"
)

func Test_test1(t *testing.T) {
	test1()
}

func Test_test2(t *testing.T) {
	test2()
}

func Test_test3(t *testing.T) {
	test3()
}

func Test_printSlice(t *testing.T) {
	var numbers = make([]int, 3, 5)
	var a [3]int
	fmt.Println(a)
	printSlice(numbers)
}
