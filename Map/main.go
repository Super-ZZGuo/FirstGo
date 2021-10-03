package main

import "fmt"

func main() {

	//不初始化map的话默认是一个nil map，且无法插入新的键值对

	var countryCapitalMap map[string]string
	fmt.Printf("map的键值元素有：%d 个\n", len(countryCapitalMap))
	println("-----不初始化map-------")

	//初始化内容
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	for index := range countryCapitalMap {
		fmt.Printf("%s 的首都是: %s \n", index, countryCapitalMap[index])
	}
	fmt.Printf("map的键值元素有：%d 个\n", len(countryCapitalMap))

	//可以通过ok-idiom获取值，来判断key/value是否存在
	//value获取的是value值，is_true获取的是true/false
	value, isTrue := countryCapitalMap["France"]
	fmt.Println(value, isTrue)

	//示范一下map的删除
	println("-----删除一个map键值对后-------")
	delete(countryCapitalMap, "France")
	for index := range countryCapitalMap {
		fmt.Printf("%s 的首都是: %s \n", index, countryCapitalMap[index])
	}
	fmt.Printf("map的键值元素有：%d 个\n", len(countryCapitalMap))
	value, isTrue = countryCapitalMap["France"]
	fmt.Println(value, isTrue)

}
