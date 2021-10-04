package main

import (
	"fmt"
)

func test() {
	defer func() {
		err := recover() //捕获异常
		if err != nil {  //说明捕获到异常
			fmt.Println("捕获到异常：", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Printf("计算结果为: %d", res)
}
func main() {
	test()
	fmt.Println("程序正常运行...")
}
