package main

import "fmt"

func main() {
	a := 100
	if a < 20 {
		fmt.Println("a小于20")
	} else {
		fmt.Println("a大于20")
	}
	fmt.Printf("a 的值为 : %d", a)
}
