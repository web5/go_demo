package structs

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 值传递
func valueTransmit() {
	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.book_id = 1
	changeBook(book1)
	fmt.Println(book1)
	changeBookByPointer(&book1)
	fmt.Println(book1)

}

func changeBook(book Books) {
	book.title = "book1_change"
}

// 通过指针参数，可改变结构体值
func changeBookByPointer(book *Books) {
	book.title = "book1_change"
}
