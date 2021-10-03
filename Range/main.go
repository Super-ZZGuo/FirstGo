package main

import "fmt"

func main() {

	//使用index索引和value值
	fmt.Println("---------使用index索引和value值---------")
	array := []int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("切片索引是：%d 的值是：%d \n", index, value)
	}

	//只使用i索引来获取数组的值
	fmt.Println("---------只使用i索引来获取数组的值---------")
	for i, _ := range array {
		fmt.Printf("切片索引是：%d 的值是：%d \n", i, array[i])
	}

	//只使用v值
	fmt.Println("---------只使用v值---------")
	for _, v := range array {
		fmt.Printf("切片的值是：%d \n", v)
	}
}
