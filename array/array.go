package array

import "fmt"

func test1() {
	var a1 [10]int
	a1[0] = 1
	fmt.Println(a1)
}

func test2() {
	var a1 = [5]float32{1000.0, 2.0, 3, 4, 50}
	fmt.Println(a1)
	fmt.Println(a1[4])
}

func test3() {
	var a1 [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		a1[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d \n", j, a1[j])
	}
}

func triangle(n int) {
	var item []int
	fmt.Printf("addr %x \n", &item)
	for i := 1; i <= n; i++ {
		item_len := len(item)
		if item_len == 0 {
			item = append(item, 1)
			fmt.Printf("addr %x \n", &item)
		} else {
			temp_s := []int{1}
			for j := 0; j < item_len-1; j++ {
				temp_s = append(temp_s, item[j]+item[j+1])
			}
			temp_s = append(temp_s, 1)
			item = temp_s
		}
		fmt.Println(item)
	}
}
