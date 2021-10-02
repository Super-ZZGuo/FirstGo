package main

import "fmt"

func main() {

	//切片的本质
	arr := [...]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}
	fmt.Printf("arr的首地址为：%p\n", &arr)
	//切片本身就是一个指针，打印地址的时候无需&地址引用
	fmt.Printf("slice的首地址为：%p\n", slice)

	fmt.Println("-------------------------")
	//测试append扩容
	var i int
	var testSlice []int
	for i = 0; i < 1100; i++ {
		testSlice = append(testSlice, i)
		fmt.Println(cap(testSlice))
	}
}
