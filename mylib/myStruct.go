package mytest

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book Books) {
	println("Print Book By Value")
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBookByAddr(book *Books) {
	println("Print Book By Addr")
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func StructTest() {
	println("=====================Struct Test=====================")

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */
	var Book3 Books /* 声明 Book3 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* book 3 描述 */
	Book3.title = "HTML 教程"
	Book3.author = "www.runoob.com"
	Book3.subject = "HTML 语言教程"
	Book3.book_id = 1234567

	/* 打印 Book1 信息 */
	println("Print Book")
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	printBook(Book2)

	/* 打印 Book3 信息 */
	printBookByAddr(&Book3)
}
