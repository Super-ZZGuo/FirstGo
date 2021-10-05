package main

import (
	"fmt"
	"runtime"
)

func init() {
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：", runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256] CPU核数
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}
func main() {
	//获取goroot目录：
	fmt.Println("GOROOT-->", runtime.GOROOT())

	//获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS) // GOOS--> darwin，mac系统

}
