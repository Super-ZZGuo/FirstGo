package main

import (
	"fmt"
	"math"
)

//使用结构体的方式能自定义/提供更多的错误信息
type radiusError struct {
	err    string
	radius float64
	area   float64
}

//实现error接口
func (e *radiusError) Error() string {
	return fmt.Sprintf("输入半径为：%0.2f，报错信息： %s，计算出面积为：%0.2f", e.radius, e.err, e.area)
}
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("面积计算失败, 输入半径 %0.2f 是一个负数", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
	//打开文件失败，返回一个非nil的错误
	//判断错误常用方法是判断返回的错误是否为nil
	//f,err := os.Open("test.txt")
	//if err != nil{
	//	fmt.Println("error_message：",err)
	//	fmt.Println("f:",f)
	//	return
	//}
	//fmt.Println(f.Name(),"打开文件成功")

	//尝试自定义一个错误信息
	radius := -100.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*radiusError); ok {
			fmt.Printf("输入半径 %0.2f 小于0", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("面积是：", area)

}
