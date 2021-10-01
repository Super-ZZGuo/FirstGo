package main

import "fmt"

func main() {
	var i int
	//快速初始化数组
	arry1 := [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	//不初始化数组长度
	arry2 := [...]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	//初始化某一索引下的元素
	arry3 := [...]float32{1: 1.0, 4: 9.0}

	fmt.Println("arry1数组如下：")
	for i = 0; i < len(arry1); i++ {
		fmt.Printf("arry1[%d] = %f\n", i, arry1[i])
	}
	fmt.Println("-----------")
	fmt.Println("arry2数组如下：")
	for i = 0; i < len(arry2); i++ {
		fmt.Printf("arry2[%d] = %f\n", i, arry2[i])
	}
	fmt.Println("-----------")
	fmt.Println("arry3数组如下：")
	for i = 0; i < len(arry3); i++ {
		fmt.Printf("arry3[%d] = %f\n", i, arry3[i])
	}
}
