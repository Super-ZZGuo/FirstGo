package main

import (
	"FirstGo/Struct/book"
	"fmt"
)

func main() {
	var Book1 book.Books
	Book1.Name = "GoGoGo"
	Book1.Author = "zzg"
	Book1.Subject = "123321"

	//直接报错
	//Book1.book_id = 123

	fmt.Printf("Book1原来的Name为：%s\n", Book1.Name)
	printBook(&Book1)

}

//测试一下结构体的指针传递
func printBook(book *book.Books) {
	//此处修改了地址的值
	book.Name = "FakeGoGoGo"
	fmt.Printf("Book1.Name = %s\n", book.Name)
	fmt.Printf("Book1.Author = %s\n", book.Author)
	fmt.Printf("Book1.Subject = %s\n", book.Subject)
}
