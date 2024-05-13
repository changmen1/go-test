package main

import "fmt"

// 结构体
type person struct {
	name    string
	age     int
	address string
	height  int
}

func main() {
	var a person
	a.name = "李瑶"
	a.age = 18
	a.address = "西安"
	a.height = 160
	fmt.Println(a)
}
