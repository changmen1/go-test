package main

import "fmt"

// 闭包 内部作用域可以拿到外部作用域的变量 但是外部作用域不能拿到内部作用域的变量
func add3() {
	ad := 5
	ad2 := func() {
		fmt.Print(ad)
	}
	ad2()
}

func add4() {
	number := 4
	fmt.Print(number)
}

func main() {
	add3()
	add4()
}
