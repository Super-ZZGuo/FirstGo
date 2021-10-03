package main

import "fmt"

func main() {

	//defer 延迟函数 后进先出-逆序运行
	//a,b,c := 1,2,3
	//defer fmt.Printf("a的值是: %d \n",a)
	//defer fmt.Printf("b的值是: %d \n",b)
	//fmt.Printf("c的值是: %d \n",c)
	//fmt.Println("----这是一条分割线-----")
	str := "string"
	fmt.Printf("%s \n", string(str))
	for _, v := range str {
		defer fmt.Printf("%c", v)
	}
}
