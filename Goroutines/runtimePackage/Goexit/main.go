package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {

		fmt.Println("开始一个goroutine")
		defer fmt.Println("开始defer。。。")

		//终止所在的协程
		runtime.Goexit()

		fmt.Println("goroutine结束")
	}()

	//区分主协程
	time.Sleep(5 * time.Second)
	fmt.Println("主goroutine即将结束")
}
